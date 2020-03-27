package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/346285234/bbs-server/pkg/database"

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
	flag.StringVar(&configPath, "config", "./configs/config.json", "setting config file path")
	flag.Parse()
	configs.LoadConfig(configPath)

	// Log.
	file, err := os.OpenFile(configs.Config.LogPath, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)

	// Setting db.
	db := database.Open("mysql", configs.Config.MySQLURL)

	categoryService := database.NewCategoryService(db)
	commentService := database.NewCommentService(db)
	favoriteService := database.NewFavoriteService(db)
	likeService := database.NewLikeService(db)
	tagService := database.NewTagService(db)
	topicService := database.NewTopicService(db)

	categoryHandler := router.NewCategoryHandler(&categoryService)
	commentHandler := router.NewCommentHandler(&commentService)
	favoriteHandler := router.NewFavoriteHandler(&favoriteService)
	likeHandler := router.NewLikeHandler(&likeService)
	tagHandler := router.NewTagHandler(&tagService)
	topicHandler := router.NewTopicHandler(&topicService)

	followService := database.NewFollowService(db)
	followHandler := router.NewFollowHandler(&followService)

	handlers := []interface{}{categoryHandler, commentHandler, favoriteHandler, likeHandler,
		tagHandler, topicHandler, followHandler}

	r := router.NewRouter(handlers)
	server := &http.Server{
		Addr:           configs.Config.Address,
		Handler:        r,
		ReadTimeout:    time.Duration(configs.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(configs.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
		TLSConfig:      nil,
	}

	//server.ListenAndServeTLS("cert.pem", "key.pem")
	server.ListenAndServe()
}
