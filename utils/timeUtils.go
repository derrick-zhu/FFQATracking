package utils

import (
	"time"
)

const (
	TimeLayout string = "2006-01-02 15:04:05"
)

// TimeIntervalSince1970 fetch time interval since 1970.1.1
func TimeIntervalSince1970() int64 {
	return time.Now().Unix()
}

// TimeFromTick convert time interval into time.Time
func TimeFromTick(tick int64) time.Time {
	return time.Unix(tick, 0)
}

// FormatedTimeFromTick format time with given format
func FormatedTimeFromTick(tick int64, format string) string {
	return TimeFromTick(tick).Format(format)
}

// StandardFormatedTimeFromTick format time with standard format
func StandardFormatedTimeFromTick(tick int64) string {
	return FormatedTimeFromTick(tick, "2006 Jan 2 Mon 15:04:05")
}
