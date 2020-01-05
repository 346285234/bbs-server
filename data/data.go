package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func OpenDB(url string) {
	var err error
	db, err = gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	models := []interface{}{&Topic{}, &TopicFavorite{}, &TopicLike{},
		&TopicCategory{}, &Comment{}, &CommentLike{}}

	for _, v := range models {
		db.AutoMigrate(v)
	}
}
