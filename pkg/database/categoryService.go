package database

import (
	"github.com/346285234/bbs-server/pkg/bbs"
	"github.com/jinzhu/gorm"
)

type CategoryService struct {
	op categoryOperation
}

func NewCategoryService(db *gorm.DB) CategoryService {
	op := newCategoryOperation(db)
	return CategoryService{op}
}

func (c *CategoryService) Categories() (categories []bbs.Category, err error) {
	return c.op.list()
}

// Operation.

type categoryOperation struct {
	db *gorm.DB
}

func newCategoryOperation(db *gorm.DB) categoryOperation {
	return categoryOperation{db}
}

func (c *categoryOperation) list() (categories []bbs.Category, err error) {
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *categoryOperation) add(category *bbs.Category) (err error) {
	return c.db.Create(category).Error
}

func (c *categoryOperation) remove(id uint) (err error) {
	return c.db.Where("id = ?", id).Delete(&bbs.Category{}).Error
}
