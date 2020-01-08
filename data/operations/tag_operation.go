package operations

import (
	"github.com/346285234/bbs-server/data"
	"github.com/346285234/bbs-server/data/models"
)

type tagOperation struct {
}

var TagO = newTagOperation()

func newTagOperation() *tagOperation {
	return &tagOperation{}
}

func (_ *tagOperation) List() (tags []models.Tag, err error) {
	if err := data.Db.Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (_ *tagOperation) add(tag models.Tag) (err error) {
	return data.Db.Create(&tag).Error
}

func (_ *tagOperation) remove(id uint) (err error) {
	return data.Db.Where("id = ?", id).Delete(&models.Tag{}).Error
}