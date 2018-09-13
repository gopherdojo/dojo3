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
	OmikuziPattern  []Fortune
}

//Draw おみくじを引く
func (o *Omikuzi) Draw() Fortune {
	return o.OmikuziPattern[o.RandomGenerater.Get()]
}

func init() {
	OmikuziPattern = []Fortune{Kyo, Kichi, Kichi, Chukichi, Chukichi, Daikichi}
}

//Draw おみくじを引く
func Draw() {

	o := Omikuzi{RandomGenerater: &random{}, OmikuziPattern: OmikuziPattern}
	s := o.Draw()
	println(s)
}

type random struct{}

func (r *random) Get() int {
	t := time.Now().UnixNano()
	rand.Seed(t)
	return rand.Intn(len(OmikuziPattern))
}
