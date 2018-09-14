package omikuji

import (
	"time"

	"github.com/gopherdojo/dojo3/kadai4/shuntaka9576/kuji"
)

type Clock interface {
	Now() time.Time
}

type Omikuji struct {
	Clock Clock
}

func (o *Omikuji) now() time.Time {
	if o.Clock == nil {
		return time.Now()
	}
	return o.Clock.Now()
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}

func (o *Omikuji) Run() (result string) {
	_, month, day := o.now().Date()
	if month == time.January && day == 1 || day == 2 || day == 3 {
		return kuji.Dikichi.PrintFortune()
	}

	result = kuji.RandomFortuneExpected(
		[]kuji.Kuji{
			kuji.Kichi,
			kuji.Chuukichi,
			kuji.Syoukichi,
			kuji.Suekichi,
			kuji.Kyou,
			kuji.Daikyou,
		},
	)
	return
}
