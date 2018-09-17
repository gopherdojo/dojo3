package dateUtil

import (
	"log"
	"testing"
	"time"
)

func TestIsNewYearsHoliday(t *testing.T) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println(err)
	}

	cases := []struct {
		year     int
		month    time.Month
		day      int
		expected bool
	}{
		{2018, 12, 31, false},
		{2019, 1, 1, true},
		{2019, 1, 2, true},
		{2019, 1, 3, true},
		{2019, 1, 4, false},
	}

	for _, c := range cases {
		timeNow = time.Date(c.year, c.month, c.day, 0, 0, 0, 0, jst)
		actual := IsNewYearsHoliday()

		if actual != c.expected {
			t.Errorf("%v is expected, but returned %v", c.expected, actual)
		}
	}
}
