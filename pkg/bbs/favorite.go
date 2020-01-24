package bbs

import "github.com/jinzhu/gorm"

type Favorite struct {
	gorm.Model
	TopicID uint
	UserID  uint
}

type FavoriteService interface {
	Mark(favorite Favorite, isMark bool) (err error)
	Check(favorite Favorite) (bool, error)
}
