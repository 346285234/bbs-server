package gorm

import (
	"github.com/346285234/bbs-server/configs"
	"github.com/346285234/bbs-server/pkg/bbs"
	"github.com/jinzhu/gorm"
)

func Open(name, url string) *gorm.DB {

	db, err := gorm.Open("mysql", configs.Config.MySQLURL)
	if err != nil {
		panic(err)
	}
	var m = []interface{}{&bbs.Topic{}, &bbs.Category{}, &bbs.Tag{}, &bbs.Favorite{}, &bbs.Like{}, &bbs.Comment{}}
	for _, v := range m {
		db.AutoMigrate(v)
	}

	return db
}
