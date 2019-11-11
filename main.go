package main

import (
	"github.com/346285234/bbs-server/router"
	"net/http"
	"time"
)

var config Configuration

func init() {
	loadConfig()
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:			config.Address,
		Handler:		mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	router.ListTopics(w, r)
}
