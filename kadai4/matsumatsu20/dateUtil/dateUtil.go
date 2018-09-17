package dateUtil

import (
	"time"
	"log"
)

func IsNewYearsHoliday() bool {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Println(err)
	}

	_, month, day := time.Now().In(jst).Date()
	if month == time.January && (day == 1 || day == 2 || day == 3) {
		return true
	}

	return false
}