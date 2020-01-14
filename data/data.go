package data

import (
	"github.com/346285234/bbs-server/data/models"
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

	for _, v := range models.Models {
		Db.AutoMigrate(v)
	}
}
