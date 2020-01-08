package operations

import (
	"github.com/346285234/bbs-server/data"
	"github.com/346285234/bbs-server/data/models"
)

type categoryOperation struct {
}

var Co = newCategoryOperation()

func newCategoryOperation() *categoryOperation {
	return &categoryOperation{}
}

func (_ *categoryOperation) List() (categories []models.Category, err error) {
	if err := data.Db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (_ *categoryOperation) add(category models.Category) (err error) {
	return data.Db.Create(&category).Error
}

func (_ *categoryOperation) remove(id uint) (err error) {
	return data.Db.Where("id = ?", id).Delete(&models.Category{}).Error
}