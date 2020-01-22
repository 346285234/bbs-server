package main

import (
	"flag"
	"net/http"
	"testing"
	"time"

	"github.com/346285234/bbs-server/pkg/gorm"

	"github.com/346285234/bbs-server/configs"
	"github.com/346285234/bbs-server/pkg/router"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
}

func main() {

	// make test init before flag.
	var _ = func() bool {
		testing.Init()
		return true
	}()

	// Load config.
	var configPath string
	flag.StringVar(&configPath, "config", "./config.json", "setting config file path")
	flag.Parse()
	configs.LoadConfig(configPath)

	// Log.
	//file, err := os.OpenFile(common.Config.LogPath, os.O_CREATE|os.O_APPEND, 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//log.SetOutput(file)
	//log.SetFormatter(&log.JSONFormatter{})
	//log.SetLevel(log.WarnLevel)

	// Setting db.
	gorm.Open("mysql", configs.Config.MySQLURL)

	r := router.NewRouter()
	server := &http.Server{
		Addr:           configs.Config.Address,
		Handler:        r,
		ReadTimeout:    time.Duration(configs.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(configs.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
		TLSConfig:      nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
