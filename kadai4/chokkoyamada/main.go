package main

import (
	"net/http"
	"time"
	"math/rand"
	"encoding/json"
	"log"
)

var timeNowFunc = time.Now
var index = func() int {
	return rand.Intn(6)
}

type Result struct {
	Omikuji string `json:"omikuji"`
	Time    string `json:"time"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var result string
	omikuji := []string{"大吉", "中吉", "小吉", "吉", "凶", "大凶"}
	index := index()
	now := timeNowFunc()

	_, m, d := now.Date()
	if m.String() == "January" && (d == 1 || d == 2 || d == 3) {
		index = 0
	}

	result = omikuji[index]

	v := Result{
		Omikuji: result,
		Time:    timeNowFunc().Format("2006-01-02 15:04:05"),
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Println("Error", err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
