package lot_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gopherdojo/dojo3/kadai4/daikurosawa/clock"

	l "github.com/gopherdojo/dojo3/kadai4/daikurosawa/lot"
)

func TestLot_Handler(t *testing.T) {
	lot := l.NewLot(nil)
	ts := httptest.NewServer(http.HandlerFunc(lot.Handler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("failed request. URL: %s, err: %s", ts.URL, err)
	}

	if res.StatusCode != 200 {
		t.Fatalf("failed error response. status_code: %v, body: %s", res.StatusCode)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("response body read error.", err)
	}

	jsonBody := new(l.ExportResponse)
	if err := json.Unmarshal(b, jsonBody); err != nil {
		t.Fatal("response body purse error.", err)
	}

	if len(jsonBody.Result) == 0 {
		t.Fatal("response body \"result\" is empty.")
	}
}

func TestNewLot_Handler_NewYear(t *testing.T) {
	lot := l.NewLot(nil)
	ts := httptest.NewServer(http.HandlerFunc(lot.Handler))
	defer ts.Close()

	cases := []struct {
		name string
		cf   clock.ClockFunc
	}{
		{
			name: "2018/01/01",
			cf: func() time.Time {
				return time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)
			},
		},
		{
			name: "2018/01/02",
			cf: func() time.Time {
				return time.Date(2018, 1, 2, 0, 0, 0, 0, time.Local)
			},
		},
		{
			name: "2018/01/03",
			cf: func() time.Time {
				return time.Date(2018, 1, 3, 0, 0, 0, 0, time.Local)
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			lot.Clock = c.cf
			res, err := http.Get(ts.URL)
			if err != nil {
				t.Errorf("failed request. URL: %s, err: %s", ts.URL, err)
			}

			if res.StatusCode != 200 {
				t.Fatalf("failed error response. status_code: %v, body: %s", res.StatusCode)
			}
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal("response body read error.", err)
			}

			jsonBody := new(l.ExportResponse)
			if err := json.Unmarshal(b, jsonBody); err != nil {
				t.Fatal("response body purse error.", err)
			}

			if jsonBody.Result != "大吉" {
				t.Fatalf("response body \"result\" is not \"大吉\". result: %s", jsonBody.Result)
			}
		})
	}
}
