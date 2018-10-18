package utils

import (
	"time"
)

const (
	// TimeLayout standard time format
	TimeLayout = "2006-01-02 15:04"
	DateLayout = "2006-01-02"
)

// TimeTickSince1970 fetch time interval since 1970.1.1
func TimeTickSince1970() int64 {
	return time.Now().Unix()
}

// TimeFromTick convert time interval into time.Time
func TimeFromTick(tick int64) time.Time {
	return time.Unix(tick, 0)
}

// TickFromDate covert date string into timeinterval
func TickFromDate(date string) int64 {

	tick, err := time.Parse(DateLayout, date)
	if err != nil {
		return -1
	}

	return tick.Unix()
}

// FormatedTimeFromTick format time with given format
func FormatedTimeFromTick(tick int64, format string) string {
	return TimeFromTick(tick).Format(format)
}

// StandardFormatedTimeFromTick format time with standard format
func StandardFormatedTimeFromTick(tick int64) string {
	return FormatedTimeFromTick(tick, TimeLayout)
}

// StandardFormatedDateFormTick format time with standard date format
func StandardFormatedDateFormTick(tick int64) string {
	return FormatedTimeFromTick(tick, DateLayout)
}
