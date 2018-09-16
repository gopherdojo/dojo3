package shimastripe

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var fortuneList []string

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	os.Exit(ret)
}

func setup() {
	for index := 0; index < int(threshold); index++ {
		fortuneList = append(fortuneList, FortuneElement(index).String())
	}
}

// Test handler. This test is flaky. (choose fortune element randomly)
func TestHandlerRandomly(t *testing.T) {
	cli := &CLI{Clock: ClockFunc(func() time.Time {
		return time.Now()
	})}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	cli.handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()

	if rw.StatusCode != http.StatusOK {
		t.Error("unexpected status code\n")
	}

	m := struct {
		Fortune string `json:"fortune"`
	}{}

	if err := json.NewDecoder(rw.Body).Decode(&m); err != nil {
		t.Error("Decode error\n")
	}

	if !contain(fortuneList, m.Fortune) {
		t.Errorf("unexpected fortune element: %v", m.Fortune)
	}
}

func TestHandlerWhenNewYear(t *testing.T) {
	cases := []struct {
		clock  Clock
		answer string
	}{
		{clock: mockClock(t, "2018/01/01"), answer: "大吉"},
		{clock: mockClock(t, "2018/01/02"), answer: "大吉"},
		{clock: mockClock(t, "2018/01/03"), answer: "大吉"},
	}

	for _, c := range cases {
		t.Run("NewYearCase", func(t *testing.T) {
			t.Helper()
			cli := &CLI{Clock: c.clock}
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			cli.handler(w, r)
			rw := w.Result()
			defer rw.Body.Close()

			if rw.StatusCode != http.StatusOK {
				t.Error("unexpected status code\n")
			}

			m := struct {
				Fortune string `json:"fortune"`
			}{}

			if err := json.NewDecoder(rw.Body).Decode(&m); err != nil {
				t.Error("Decode error\n")
			}

			if m.Fortune != c.answer {
				t.Errorf("unexpected fortune element: %v, expected: %v", m.Fortune, c.answer)
			}
		})
	}
}

func contain(list []string, elm string) bool {
	for _, l := range list {
		if l == elm {
			return true
		}
	}
	return false
}

func mockClock(t *testing.T, v string) Clock {
	t.Helper()
	now, err := time.Parse("2006/01/02", v)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	return ClockFunc(func() time.Time {
		return now
	})
}
