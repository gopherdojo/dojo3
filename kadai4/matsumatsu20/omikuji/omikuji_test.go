package omikuji

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func TestHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(Handler))
	defer ts.Close()

	cases := []struct{
		luck              string
		isNewYearsHoliday bool
		expected          response
	}{
		{daikichi, false, response{200, "大吉"}},
		{kyo, false, response{200, "凶"}},
		{shokichi, true, response{200, "大吉"}},
	}

	for _, c := range cases {
		fetchKujiFunc = func() string {
			return c.luck
		}
		isNewYearsHolidayFunc = func() bool {
			return c.isNewYearsHoliday
		}

		res, err := http.Get(ts.URL)

		if err != nil {
			t.Fatal(err)
		}

		body, err  := ioutil.ReadAll(res.Body)

		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != 200 {
			t.Errorf("expedted status %v, but returned %v", c.expected, 200)
		}

		var actual response
		json.Unmarshal(body, &actual)

		if actual != c.expected {
			t.Errorf("expedted %v, but returned %v", c.expected, actual)
		}
	}
}
