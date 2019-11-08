package common

import "net/http"

func CheckUser(r *http.Request) bool {
	return true
}

func Intro(from string) string {
	return from[:1]
}