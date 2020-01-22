package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Value  string
	topics []Topic
}

type CategoryService interface {
	Categories() (categories []Category, err error)
}
