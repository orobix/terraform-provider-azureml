package provider

import (
	"hash/fnv"
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

func hash(s string) (uint32, error) {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return 0, err
	}
	return h.Sum32(), nil
}
