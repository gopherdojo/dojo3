package omikuzi

import (
	"testing"
)

func TestDraw(t *testing.T) {
	cases := []struct {
		name   string
		index  int
		result Fortune
	}{
		{name: "normal_0", index: 0, result: Kyo},
		{name: "normal_1", index: 1, result: Kichi},
		{name: "normal_2", index: 2, result: Kichi},
		{name: "normal_3", index: 3, result: Chukichi},
		{name: "normal_4", index: 4, result: Chukichi},
		{name: "normal_5", index: 5, result: Daikichi},
	}
	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			o := Omikuzi{RandomGenerater: &TestRandom{index: v.index}, OmikuziPattern: OmikuziPattern}
			result := o.Draw()
			if result != v.result {
				t.Errorf("expected %v but %v ", v.result, result)
			}
		})
	}
}

func getTestRandomGenerater(index int) {

}

type TestRandom struct {
	index int
}

func (r *TestRandom) Get() int {
	return r.index
}
