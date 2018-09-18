package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var fortunes = []string{"大吉", "吉", "中吉", "小吉", "末吉", "凶", "大凶"}

type Omikuji struct {
	Date   time.Time `json:"date"`
	Result string    `json:"result"`
}

func NewOmikuji(date time.Time) *Omikuji {
	o := Omikuji{Date: date}
	return &o
}

type Adapter func(http.Handler) http.Handler

func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func SetHeader() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			h.ServeHTTP(w, r)
		})
	}
}

func (o *Omikuji) handle(w http.ResponseWriter, r *http.Request) {
	if o.Date.Month() == 1 && (o.Date.Day() >= 1 && o.Date.Day() <= 3) {
		o.Result = "大吉"
	} else {
		o.Result = fortunes[rand.Int()%len(fortunes)]
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(o); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, buf.String())
}

func main() {
	o := NewOmikuji(time.Now())
	handler := http.HandlerFunc(o.handle)
	http.Handle("/", Adapt(handler, SetHeader()))
	http.ListenAndServe(":8080", nil)
}
