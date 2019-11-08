package main

import (
	"fmt"
	"net/http"
)

func main() {
	files := http.FileServer(http.Dir("/public/"))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	address := ":8000"
	http.HandleFunc("/", index)

	http.ListenAndServe(address, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
