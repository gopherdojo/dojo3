package omikuzi

import (
	"log"
	"math/rand"
	"time"
)

//Fortune おみくじ結果
type Fortune string

//RandomGenerater ランダム生成器
type RandomGenerater interface {
	Get() int
}

//OmikuziPattern おみくじ結果配列
var OmikuziPattern []Fortune

const (
	//Daikichi 大吉
	Daikichi Fortune = "大吉"

	//Chukichi 中吉
	Chukichi Fortune = "中吉"

	//Kichi 吉
	Kichi Fortune = "吉"

	//Kyo 凶
	Kyo Fortune = "凶"
)

//TimeLayout 日付レイアウト
const TimeLayout = "2006-01-02"

//Omikuzi おみくじ構造体
type Omikuzi struct {
	RandomGenerater RandomGenerater
	Time            time.Time
	OmikuziPattern  []Fortune
}

//Draw おみくじを引く
func (o *Omikuzi) Draw() Fortune {
	m := int(o.Time.Month())
	d := o.Time.Day()

	var r Fortune
	if isShogatsu(m, d) {
		r = Daikichi
	} else {
		r = o.OmikuziPattern[o.RandomGenerater.Get()]
	}
	return r
}

func isShogatsu(m int, d int) bool {
	if m == 1 && d >= 1 && d <= 3 {
		return true
	}
	return false
}

func init() {
	OmikuziPattern = []Fortune{
		Kyo,
		Kichi,
		Kichi,
		Chukichi,
		Chukichi,
		Daikichi,
	}
}

//Draw おみくじを引く
func Draw() string {

	now := time.Now()

	//XXX 1日で結果を固定するように当日の00:00を取得するためフォーマット化し、
	//またtime.Timeに戻してている。
	//別のいいやり方がないか検討していた。
	timeFormated := now.Format(TimeLayout)
	targetTime, err := time.Parse(TimeLayout, timeFormated)
	if err != nil {
		log.Println("error:", err)
	}

	o := Omikuzi{
		RandomGenerater: &random{TargetTime: targetTime},
		Time:            targetTime,
		OmikuziPattern:  OmikuziPattern,
	}
	return string(o.Draw())
}

//DrawByDate 日付でおみくじを引く
func DrawByDate(targetDate string) string {

	targetTime, err := time.Parse(TimeLayout, targetDate)
	if err != nil {
		log.Println("error:", err)
	}

	o := Omikuzi{
		RandomGenerater: &random{TargetTime: targetTime},
		Time:            targetTime,
		OmikuziPattern:  OmikuziPattern,
	}
	return string(o.Draw())
}

type random struct {
	TargetTime time.Time
}

func (r *random) Get() int {
	rand.Seed(r.TargetTime.Unix())
	return rand.Intn(len(OmikuziPattern))
}
