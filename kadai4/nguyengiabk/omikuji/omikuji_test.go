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

var testGetResultFixtures = map[string]struct {
	clock     omikuji.Clock
	randomize omikuji.Randomize
	result    omikuji.Fortune
}{
	"Test 1/1": {
		omikuji.ClockFunc(func() time.Time {
			return time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)
		}),
		nil,
		omikuji.Daikichi,
	},
	"Test 1/2": {
		omikuji.ClockFunc(func() time.Time {
			return time.Date(2018, 1, 2, 0, 0, 0, 0, time.Local)
		}),
		nil,
		omikuji.Daikichi,
	},
	"Test 1/3": {
		omikuji.ClockFunc(func() time.Time {
			return time.Date(2018, 1, 3, 0, 0, 0, 0, time.Local)
		}),
		nil,
		omikuji.Daikichi,
	},
	"Test 1/4 and Kyou": {
		omikuji.ClockFunc(func() time.Time {
			return time.Date(2018, 1, 4, 0, 0, 0, 0, time.Local)
		}),
		omikuji.RandomizeFunc(func(max int) int {
			return 4
		}),
		omikuji.Kyou,
	},
	"Test Daikichi not new year": {
		omikuji.ClockFunc(func() time.Time {
			return time.Date(2018, 7, 1, 0, 0, 0, 0, time.Local)
		}),
		omikuji.RandomizeFunc(func(max int) int {
			return 0
		}),
		omikuji.Daikichi,
	},
	"Test Chukichi": {
		nil,
		omikuji.RandomizeFunc(func(max int) int {
			return 1
		}),
		omikuji.Chukichi,
	},
	"Test Shokichi": {
		nil,
		omikuji.RandomizeFunc(func(max int) int {
			return 2
		}),
		omikuji.Shokichi,
	},
	"Test Shokyou": {
		nil,
		omikuji.RandomizeFunc(func(max int) int {
			return 5
		}),
		omikuji.Shokyou,
	},
	"Test Daikyou": {
		nil,
		omikuji.RandomizeFunc(func(max int) int {
			return 6
		}),
		omikuji.Daikyou,
	},
}

func TestGetResult(t *testing.T) {
	for name, tc := range testGetResultFixtures {
		t.Run(name, func(t *testing.T) {
			o := omikuji.Omikuji{Clock: tc.clock, Randomize: tc.randomize}
			result := o.GetResult()
			if result != tc.result {
				t.Errorf("GetResult() return wrong result, actual = %v, expected = %v", result, tc.result)
			}
		})
	}
}
