package main

import (
	"flag"
	"net/http"
	"math/rand"
	"time"
	"encoding/json"
	"log"
)

var (
	port = flag.String("p", "8080", "resten port")

	luck = map[int]string{
		0: "大吉",
		1: "吉",
		2: "中吉",
		3: "小吉",
		4: "末吉",
		5: "凶",
		6: "大凶",
	}

)

type Response struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}

type Holiday struct {
	Month time.Month
	Day   int
}

type Holidays struct {
	Days []Holiday
}

func main() {
	http.HandleFunc("/kuji", omikuji)

	http.ListenAndServe(":" + *port, nil)
}

func omikuji(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	i := rand.Intn(len(luck))

	// TODO: マジックナンバーを消したい
	if isNewYearsHolidays() {
		i = 0
	}

	res := &Response{Status: 200, Result: luck[i]}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(res); err != nil {log.Fatal(err)}
}

func isNewYearsHolidays() bool {
	newYearsHolidays := Holidays{
		[]Holiday{
			{time.Month(1), 1},
			{time.Month(1), 2},
			{time.Month(1), 3},
		},
	}

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println(err)
	}
	if newYearsHolidays.include(time.Now().In(jst)) {
		return true
	}

	return false
}

func (h *Holidays) include(date time.Time) bool {
	_, month, day := date.Date()

	for _, holiday := range h.Days {
		if holiday.Month.String() == month.String() && holiday.Day == day {
			return true
		}
	}

	return true
}
