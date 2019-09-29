package stringUtil

import (
	"strconv"
)

/**
 * check if string slice contains string value
 */
func InSlice (a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}


func ToInt64 (value string) int64 {

	intValue, err := strconv.ParseInt(value, 10, 64)

	if err == nil {
		return intValue
	}

	return 0
}


func ToInt (value string) int {

	intValue, err := strconv.Atoi(value)

	if err == nil {
		return intValue
	}

	return 0
}