package data

import (
	models2 "github.com/346285234/bbs-server/data/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func OpenDB(url string) {
	var err error
	Db, err = gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	models := []interface{}{&models2.Topic{}, &models2.TopicFavorite{}, &models2.TopicLike{},
		&models2.TopicCategory{}, &models2.Comment{}, &models2.CommentLike{}}

	for _, v := range models {
		Db.AutoMigrate(v)
	}
}
