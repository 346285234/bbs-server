package gorm

import (
	"github.com/346285234/bbs-server/pkg/models"
	"github.com/jinzhu/gorm"
)

type CategoryService struct {
	op categoryOperation
}

func NewCategoryService(db *gorm.DB) CategoryService {
	op := newCategoryOperation(db)
	return CategoryService{op}
}

func (c *CategoryService) Categories() (categories []models.Category, err error) {
	return c.op.list()
}

// Operation.

type categoryOperation struct {
	db *gorm.DB
}

func newCategoryOperation(db *gorm.DB) categoryOperation {
	return categoryOperation{db}
}

func (c *categoryOperation) list() (categories []models.Category, err error) {
	if err := c.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *categoryOperation) add(category *models.Category) (err error) {
	return c.db.Create(category).Error
}

func (c *categoryOperation) remove(id uint) (err error) {
	return c.db.Where("id = ?", id).Delete(&models.Category{}).Error
}
