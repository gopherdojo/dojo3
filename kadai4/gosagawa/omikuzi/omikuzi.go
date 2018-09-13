package omikuzi

import (
	"math/rand"
	"time"
)

type Fortune string
type RandomGenerater interface {
	Get() int
}

var OmikuziPattern []Fortune

const (
	Daikichi Fortune = "大吉"
	Chukichi Fortune = "中吉"
	Kichi    Fortune = "吉"
	Kyo      Fortune = "凶"
)

type Omikuzi struct {
	RandomGenerater RandomGenerater
	OmikuziPattern  []Fortune
}

func (o *Omikuzi) Draw() Fortune {
	return o.OmikuziPattern[o.RandomGenerater.Get()]
}

func init() {
	OmikuziPattern = []Fortune{Kyo, Kichi, Kichi, Chukichi, Chukichi, Daikichi}
}

func Draw() {

	o := Omikuzi{RandomGenerater: &Random{}, OmikuziPattern: OmikuziPattern}
	s := o.Draw()
	println(s)
}

type Random struct{}

func (r *Random) Get() int {
	t := time.Now().UnixNano()
	rand.Seed(t)
	return rand.Intn(len(OmikuziPattern))
}
