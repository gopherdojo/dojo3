package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestFortuneHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/fortune", nil)

	tr := &Timer{Now: time.Date(2018, time.January, 2, 0, 0, 0, 0, time.UTC)}
	tr.FortuneHandler(w, r)

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
	if s := string(b); !strings.Contains(s, expected) {
		t.Fatalf("unexpected response: %s", s)
	}
}
