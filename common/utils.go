package common

import (
	"net/http"
	"strconv"
)

func CheckUser(r *http.Request) bool {
	return true
}

func Intro(from string) string {
	return from[:1]
}

func StrToInt(from string) int {
	result, error := strconv.Atoi(from)
	if error != nil {
		result = 0
	}

	return result
}
