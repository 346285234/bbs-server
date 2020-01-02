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
	r := router.NewRouter()
	server := &http.Server{
		Addr:			config.Address,
		Handler:		r,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
		TLSConfig: nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}