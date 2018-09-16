package omikuji_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gopherdojo/dojo3/kadai4/nguyengiabk/omikuji"
)

func Example() {
	o := omikuji.Omikuji{}
	fmt.Println(o.GetResult())
}

type randomize struct {
	useMock    bool
	randResult int
}

var testGetResultFixtures = map[string]struct {
	date      string
	randomize randomize
	result    omikuji.Fortune
}{
	"Test 1/1":                   {"2018/01/01", randomize{false, 0}, omikuji.Daikichi},
	"Test 1/2":                   {"2018/01/02", randomize{false, 0}, omikuji.Daikichi},
	"Test 1/3":                   {"2018/01/03", randomize{false, 0}, omikuji.Daikichi},
	"Test 1/4 and Kyou":          {"2018/01/04", randomize{true, 4}, omikuji.Kyou},
	"Test Daikichi not new year": {"2018/07/01", randomize{true, 0}, omikuji.Daikichi},
	"Test Chukichi":              {"2018/08/02", randomize{true, 1}, omikuji.Chukichi},
	"Test Shokichi":              {"2018/09/10", randomize{true, 2}, omikuji.Shokichi},
	"Test Shokyou":               {"2018/10/11", randomize{true, 5}, omikuji.Shokyou},
	"Test Daikyou":               {"2018/03/21", randomize{true, 6}, omikuji.Daikyou},
}

func TestGetResult(t *testing.T) {
	for name, tc := range testGetResultFixtures {
		tc := tc
		t.Run(name, func(t *testing.T) {
			clock := mockClock(t, tc.date)
			var randomize omikuji.Randomize
			if tc.randomize.useMock {
				randomize = mockRandomize(t, tc.randomize.randResult)
			}
			o := omikuji.Omikuji{Clock: clock, Randomize: randomize}
			result := o.GetResult()
			if result != tc.result {
				t.Errorf("GetResult() return wrong result, actual = %v, expected = %v", result, tc.result)
			}
		})
	}
}

func mockClock(t *testing.T, v string) omikuji.Clock {
	t.Helper()
	now, err := time.Parse("2006/01/02", v)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	return omikuji.ClockFunc(func() time.Time {
		return now
	})
}

func mockRandomize(t *testing.T, v int) omikuji.Randomize {
	t.Helper()
	return omikuji.RandomizeFunc(func(max int) int {
		return v
	})
}
