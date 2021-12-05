package provider

import (
	"regexp"
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func stringIsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func stringIsOnlyLettersAndDigits(s string) bool {
	res, _ := regexp.MatchString("^[A-Za-z0-9]*$", s)
	return res
}
