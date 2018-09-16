package omikuji

import (
	"math/rand"
	"time"
)

// Fortune represents omikuji result
type Fortune string

const (
	// Daikichi represents 大吉 result
	Daikichi Fortune = "大吉"

	// Chukichi represents 中吉 result
	Chukichi Fortune = "中吉"

	// Shokichi represents 小吉 result
	Shokichi Fortune = "小吉"

	// Kichi represents 吉 result
	Kichi Fortune = "吉"

	// Kyou represents 凶 result
	Kyou Fortune = "凶"

	// Shokyou represents 小凶 result
	Shokyou Fortune = "小凶"

	// Daikyou represents 大凶 result
	Daikyou Fortune = "大凶"
)

var omikujiValues = []Fortune{Daikichi, Chukichi, Shokichi, Kichi, Kyou, Shokyou, Daikyou}

// Clock defines types that can return current time
type Clock interface {
	Now() time.Time
}

// ClockFunc returns current time, we use this type for testing
type ClockFunc func() time.Time

// Now returns current time by calling CockFunc itself
func (f ClockFunc) Now() time.Time {
	return f()
}

// Randomize defines types that can return ramdom integer
type Randomize interface {
	Rand(max int) int
}

// RandomizeFunc returns randome integer, we use this type for testing
type RandomizeFunc func(max int) int

// Rand returns random integer by calling RandomizeFunc itself
func (f RandomizeFunc) Rand(max int) int {
	return f(max)
}

// Omikuji is used to get omikuji result based on current time
type Omikuji struct {
	Clock     Clock
	Randomize Randomize
}

func (o *Omikuji) now() time.Time {
	if o.Clock == nil {
		return time.Now()
	}
	return o.Clock.Now()
}

func (o *Omikuji) rand(max int) int {
	if o.Randomize == nil {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(max)
	}
	return o.Randomize.Rand(max)
}

// GetResult return randomly selected omikuji value
func (o *Omikuji) GetResult() Fortune {
	_, m, d := o.now().Date()
	switch {
	case m == time.January && d <= 3:
		return Daikichi
	default:
		return omikujiValues[o.rand(len(omikujiValues))]
	}
}
