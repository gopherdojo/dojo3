package omikuzi

import (
	"math/rand"
	"time"
)

type Fortune string

var OmikuziPattern []Fortune

const (
	Daikichi Fortune = "大吉"
	Chukichi Fortune = "中吉"
	Kichi    Fortune = "吉"
	Kyo      Fortune = "凶"
)

func init() {
	OmikuziPattern = []Fortune{Kyo, Kichi, Kichi, Chukichi, Chukichi, Daikichi}
}

func Draw() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(len(OmikuziPattern))
	println(OmikuziPattern[s])
}
