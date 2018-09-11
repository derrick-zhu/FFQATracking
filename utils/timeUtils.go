package utils

import (
	"time"
)

// TimeIntervalSince1970 fetch time interval since 1970.1.1
func TimeIntervalSince1970() int64 {
	return time.Now().Unix()
}

// TimeFromTick convert time interval into time.Time
func TimeFromTick(tick int64) time.Time {
	return time.Unix(tick, 0)
}
