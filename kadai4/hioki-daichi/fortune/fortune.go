/*
Package fortune is a package that manages processing around fortune.
*/
package fortune

import (
	"math/rand"
)

// Fortune means 運勢
type Fortune string

const (
	// Daikichi means "大吉"
	Daikichi Fortune = "大吉"

	// Chukichi means "中吉"
	Chukichi Fortune = "中吉"

	// Shokichi means "小吉"
	Shokichi Fortune = "小吉"

	// Kichi means "吉"
	Kichi Fortune = "吉"

	// Suekichi means "末吉"
	Suekichi Fortune = "末吉"

	// Kyo means "凶"
	Kyo Fortune = "凶"

	// Daikyo means "大凶"
	Daikyo Fortune = "大凶"
)

// DrawFortune draws a fortune.
func DrawFortune() Fortune {
	fs := AllFortunes()
	return fs[rand.Intn(len(fs))]
}

// AllFortunes returns all fortunes.
func AllFortunes() []Fortune {
	return []Fortune{Daikichi, Chukichi, Shokichi, Kichi, Suekichi, Kyo, Daikyo}
}
