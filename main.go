package main

import (
	"flag"
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data"
	"github.com/346285234/bbs-server/router"
	"net/http"
	"time"
)

func init() {
	// load config.
	var configPath string
	flag.StringVar(&configPath, "config", "./config.json", "setting config file path")
	flag.Parse()
	common.LoadConfig(configPath)

	data.OpenDB(common.Config.MySQLURL)
}

func main() {
	r := router.NewRouter()
	server := &http.Server{
		Addr:           common.Config.Address,
		Handler:        r,
		ReadTimeout:    time.Duration(common.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(common.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
		TLSConfig:      nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
