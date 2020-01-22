package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	UserID uint
	Value  string
}

type TagService interface {
	Tags() (categories []*Tag, err error)
}
