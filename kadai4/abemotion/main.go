package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Fortune struct {
	Result string `json:"result"`
}

type Timer struct {
	Now time.Time
}

func main() {
	t := &Timer{Now: time.Now()}
	http.HandleFunc("/fortune", t.FortuneHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (t *Timer) FortuneHandler(w http.ResponseWriter, r *http.Request) {
	fs := map[int]string{
		0: "大吉",
		1: "中吉",
		2: "小吉",
		3: "吉",
		4: "末吉",
		5: "凶",
		6: "大凶",
	}
	rf := fs[rand.Intn(7)]

	if t.Now.Month().String() == "January" && arrayContains([]int{1, 2, 3}, t.Now.Day()) {
		rf = "大吉"
	}

	f := &Fortune{Result: rf}
	enc := json.NewEncoder(w)
	if err := enc.Encode(f); err != nil {
		log.Fatal(err)
	}
}

func arrayContains(arr []int, str int) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
