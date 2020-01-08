package main

import (
	"flag"
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data"
	"github.com/346285234/bbs-server/router"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"os"
	"testing"
	"time"
)

func init() {

	// make test init before flag.
	var _ = func() bool {
		testing.Init()
		return true
	}()

	// load config.
	var configPath string
	flag.StringVar(&configPath, "config", "./config.json", "setting config file path")
	flag.Parse()
	common.LoadConfig(configPath)

	// log.
	file, err := os.OpenFile(common.Config.LogPath, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)

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
