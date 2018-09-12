package datehelper

import (
	"regexp"
	"testing"
	"time"
)

func TestDatehelper_IsDuringTheNewYear(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("err %s", err)
	}

	cases := map[string]struct {
		year     int
		month    time.Month
		day      int
		expected bool
	}{
		"2018-12-31": {year: 2018, month: time.December, day: 31, expected: false},
		"2019-01-01": {year: 2019, month: time.January, day: 1, expected: true},
		"2019-01-02": {year: 2019, month: time.January, day: 2, expected: true},
		"2019-01-03": {year: 2019, month: time.January, day: 3, expected: true},
		"2019-01-04": {year: 2019, month: time.January, day: 4, expected: false},
	}

	for n, c := range cases {
		c := c
		t.Run(n, func(t *testing.T) {
			nowFunc = func() time.Time {
				return time.Date(c.year, c.month, c.day, 0, 0, 0, 0, loc)
			}

			expected := c.expected
			actual := IsDuringTheNewYear()
			if actual != expected {
				t.Errorf(`expected: "%t" actual: "%t"`, expected, actual)
			}
		})
	}
}

func TestDatehelper_IsDuringTheNewYear_Panic(t *testing.T) {
	loadLocationFunc = func(name string) (*time.Location, error) {
		return time.LoadLocation("Nonexistent/Location")
	}

	defer func() {
		err := recover()
		if err == nil {
			t.Fatal("did not panic")
		}
		expected := "cannot find Nonexistent/Location in zip file "
		actual := err.(error).Error()
		if !regexp.MustCompile(expected).MatchString(actual) {
			t.Errorf(`unmatched error: expected: "%s" actual: "%s"`, expected, actual)
		}
	}()
	IsDuringTheNewYear()
}
