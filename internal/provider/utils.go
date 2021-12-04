package provider

import "strings"

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
