package gorm

import (
	"github.com/346285234/bbs-server/configs"
	"github.com/346285234/bbs-server/pkg/models"
	"github.com/jinzhu/gorm"
)

func Open(name, url string) *gorm.DB {

	db, err := gorm.Open("mysql", configs.Config.MySQLURL)
	if err != nil {
		panic(err)
	}
	var m = []interface{}{&models.Topic{}, &models.Category{}, &models.Tag{}, &models.Favorite{}, &models.Like{}, &models.Comment{}}
	for _, v := range m {
		db.AutoMigrate(v)
	}
}
