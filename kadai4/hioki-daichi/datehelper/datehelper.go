/*
Package datehelper is a collection of convenient functions for manipulating dates.
*/
package datehelper

import "time"

// for testing
var nowFunc = time.Now
var loadLocationFunc = time.LoadLocation

// IsDuringTheNewYear returns whether the current date is the New Year or not.
func IsDuringTheNewYear() bool {
	loc, err := loadLocationFunc("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	_, month, day := nowFunc().In(loc).Date()
	if month == time.January && (day == 1 || day == 2 || day == 3) {
		return true
	}
	return false
}
