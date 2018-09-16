package omikuji

import (
	"math/rand"
	"time"
)

var defaultLots = []string{
	"大吉",
	"吉",
	"中吉",
	"小吉",
	"凶",
	"大凶",
}

type Omikuji struct {
	rand *rand.Rand
}

func NewOmikuji() Omikuji {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return Omikuji{
		rand: r,
	}
}

func (o *Omikuji) Do(lots []string) string {
	l := len(lots)
	if l != 0 {
		return lots[rand.Intn(l)]
	}

	return defaultLots[rand.Intn(len(defaultLots))]
}
