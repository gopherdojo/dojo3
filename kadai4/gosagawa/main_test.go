package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	cases := []struct {
		name          string
		date          string
		isResultFixed bool
		result        string
	}{
		{name: "normal", date: ""},
		{name: "shogatsu_12/31", date: "2018-12-31"},
		{name: "shogatsu_1/1", date: "2019-01-01", isResultFixed: true, result: "大吉"},
		{name: "shogatsu_1/2", date: "2019-01-02", isResultFixed: true, result: "大吉"},
		{name: "shogatsu_1/3", date: "2019-01-03", isResultFixed: true, result: "大吉"},
		{name: "shogatsu_1/4", date: "2019-01-04"},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/?date="+v.date, nil)
			handler(w, r)
			rw := w.Result()
			defer rw.Body.Close()
			if rw.StatusCode != http.StatusOK {
				t.Fatal("unexpected status code")
			}
			var or OmikuziResult
			dec := json.NewDecoder(rw.Body)
			if err := dec.Decode(&or); err != nil {
				t.Fatal(err)
			}

			if v.isResultFixed && or.Result != v.result {
				t.Errorf("expected %v but %v ", v.result, or.Result)
			}
		})
	}
}
