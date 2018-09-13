package omikuzi

import (
	"log"
	"testing"
	"time"
)

func TestDraw(t *testing.T) {
	cases := []struct {
		name   string
		date   string
		index  int
		result Fortune
	}{
		{name: "normal_0", date: "2018-09-01", index: 0, result: Kyo},
		{name: "normal_1", date: "2018-09-01", index: 1, result: Kichi},
		{name: "normal_2", date: "2018-09-01", index: 2, result: Kichi},
		{name: "normal_3", date: "2018-09-01", index: 3, result: Chukichi},
		{name: "normal_4", date: "2018-09-01", index: 4, result: Chukichi},
		{name: "normal_5", date: "2018-09-01", index: 5, result: Daikichi},
		{name: "shogatsu_12/31", date: "2018-12-31", index: 0, result: Kyo},
		{name: "shogatsu_1/1", date: "2019-01-01", index: 0, result: Daikichi},
		{name: "shogatsu_1/2", date: "2019-01-02", index: 0, result: Daikichi},
		{name: "shogatsu_1/3", date: "2019-01-03", index: 0, result: Daikichi},
		{name: "shogatsu_1/4", date: "2019-01-04", index: 0, result: Kyo},
	}
	layout := "2006-01-02"
	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			time, err := time.Parse(layout, v.date)
			if err != nil {
				log.Fatal(err)
			}
			o := Omikuzi{
				RandomGenerater: &TestRandom{index: v.index},
				Time:            time,
				OmikuziPattern:  OmikuziPattern,
			}
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
