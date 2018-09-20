package utils

import (
	"time"
)

const (
	// TimeLayout standard time format
	TimeLayout = "2006-01-02 15:04"
)

// TimeTickSince1970 fetch time interval since 1970.1.1
func TimeTickSince1970() int64 {
	return int64(float32(time.Now().UnixNano()) * 0.000000001)
}

// TimeFromTick convert time interval into time.Time
func TimeFromTick(tick int64) time.Time {

	return time.Unix(tick, int64(float32(tick)*0.001))
}

// FormatedTimeFromTick format time with given format
func FormatedTimeFromTick(tick int64, format string) string {
	return TimeFromTick(tick).Format(format)
}

// StandardFormatedTimeFromTick format time with standard format
func StandardFormatedTimeFromTick(tick int64) string {
	return FormatedTimeFromTick(tick, TimeLayout)
}
