package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:346285234@/bbs")
	if err != nil {
		panic(err)
	}

	models := []interface{}{&Topic{}, &TopicFavorite{}, &TopicLike{},
		&Category{}, &Comment{}, &CommentLike{}}

	for _, v := range models {
		Db.AutoMigrate(v)
	}
}