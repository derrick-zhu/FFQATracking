package utils

import "regexp"

const (
	// RegexEmail email regex expression
	RegexEmail = "[\\w!#$%&'*+/=?^_`{|}~-]+(?:\\.[\\w!#$%&'*+/=?^_`{|}~-]+)*@(?:[\\w](?:[\\w-]*[\\w])?\\.)+[\\w](?:[\\w-]*[\\w])?"
)

// MatchRegex check does origin string fits regex expression
func MatchRegex(reg, origin string) bool {
	bMatch, err := regexp.MatchString(reg, origin)
	if err != nil {
		return false
	}
	return bMatch
}

// MatchRegexEmail checking email format
func MatchRegexEmail(origin string) bool {
	return MatchRegex(RegexEmail, origin)
}
