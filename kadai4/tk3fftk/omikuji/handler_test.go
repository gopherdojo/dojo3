package omikuji

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestIsNewYear(t *testing.T) {
	cases := map[string]struct {
		month    int
		day      int
		expected bool
	}{
		"December31st": {
			12,
			31,
			false,
		},
		"January1st": {
			1,
			1,
			true,
		},
		"January2nd": {
			1, 2, true,
		},
		"January3rd": {
			1, 3, true,
		},
		"January4th": {
			1, 4, false,
		},
	}

	for k, v := range cases {
		k := k
		v := v

		t.Run(k, func(t *testing.T) {
			t.Helper()
			handler := New(
				ClockFunc(func() time.Time {
					return time.Date(2018, time.Month(v.month), v.day, 6, 0, 0, 0, time.Local)
				}),
			)

			if handler.isNewYear() != v.expected {
				t.Errorf("expected=%v, actual=%v", v.expected, !v.expected)
			}
		})
	}
}

func TestCreateResponseJSON(t *testing.T) {
	handler := New(nil)
	result := "test"
	expected := "{\"result\":\"test\"}\n"
	json, err := handler.createResponseJSON(result)
	if err != nil {
		t.Fatal("should not come here")
	}
	if strings.Compare(json, expected) != 0 {
		t.Errorf("expected='%v', actual='%v'", expected, json)
	}
}

func TestHandler(t *testing.T) {
	handler := New(ClockFunc(func() time.Time {
		return time.Date(2018, 1, 1, 6, 0, 0, 0, time.Local)
	}))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	handler.handlerFunc(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Fatal("unexpected status code")
	}

	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		t.Fatal("unexpected error")
	}
	expected := "{\"result\":\"大吉\"}\n"

	if s := string(b); s != expected {
		t.Errorf("expected='%v', actual='%v'", expected, s)
	}
}
