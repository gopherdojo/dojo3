package omikuzi

import (
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
	} else {
		return false
	}
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

	o := Omikuzi{
		RandomGenerater: &random{},
		Time:            time.Now(),
		OmikuziPattern:  OmikuziPattern,
	}
	return string(o.Draw())
}

type random struct{}

func (r *random) Get() int {
	t := time.Now().UnixNano()
	rand.Seed(t)
	return rand.Intn(len(OmikuziPattern))
}
