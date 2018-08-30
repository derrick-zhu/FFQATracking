package utils

import "strconv"

// I64toa covert int64 -> string
func I64toa(v int64) string {
	return strconv.FormatInt(v, 10)
}

// Atoi64 string -> int64
func Atoi64(v string) int64 {
	n, err := strconv.ParseInt(v, 10, 64)
	if err == nil {
		return n
	}
	return 0
}
