package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:346285234@/bbs")
	if err != nil {
		panic(err)
	}

	models := []interface{}{&Topic{}, &TopicFavorite{}, &TopicLike{},
		&Category{}, &Comment{}, &CommentLike{}}

	for _, v := range models {
		db.AutoMigrate(v)
	}
}