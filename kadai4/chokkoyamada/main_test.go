package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"math/rand"
)

func setIndex() {
	index = func() int { return 1 }
}

func setNewYearTime() {
	timeNowFunc = func() time.Time {
		t, _ := time.Parse("2006-01-02", "2018-01-01")
		return t
	}
}

func setIndexBack() {
	index = func() int {
		return rand.Intn(6)
	}
}

func setOtherTime() {
	timeNowFunc = func() time.Time {
		t, _ := time.Parse("2006-01-02", "2018-01-04")
		return t
	}
}

func TestMainNormal(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	setIndex()
	setOtherTime()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	Handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
	const expected = "中吉"
	res := new(Result)
	if err := json.Unmarshal(b, res); err != nil {
		t.Fatal("json unmarshall error")
	}
	if res.Omikuji != expected {
		t.Fatalf("result different: %v, %v", res.Omikuji, res.Time)
	}
}

func TestMainNewYear(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	setIndexBack()
	setNewYearTime()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	Handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
	const expected = "大吉"
	res := new(Result)
	if err := json.Unmarshal(b, res); err != nil {
		t.Fatal("json unmarshal error")
	}
	if res.Omikuji != expected {
		t.Fatalf("result different: %v, %v", res.Omikuji, res.Time)
	}
}
