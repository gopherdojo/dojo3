package kuji

import (
	"math/rand"
)

type Kuji int

const (
	Dikichi   Kuji = 1
	Kichi     Kuji = 2
	Chuukichi Kuji = 3
	Syoukichi Kuji = 4
	Suekichi  Kuji = 5
	Kyou      Kuji = 6
	Daikyou   Kuji = 7
)

func (k Kuji) PrintFortune() string {
	switch k {
	case 1:
		return "大吉"
	case 2:
		return "吉"
	case 3:
		return "中吉"
	case 4:
		return "小吉"
	case 5:
		return "末吉"
	case 6:
		return "凶"
	case 7:
		return "大凶"
	default:
		return "Error"
	}
}

func RandomFortuneExpected(kujes []Kuji) string {
	var fortunes []Kuji
	for _, kuji := range kujes {
		fortunes = append(fortunes, kuji)
	}
	return fortunes[rand.Intn(len(fortunes))].PrintFortune()
}
