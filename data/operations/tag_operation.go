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

func (_ *tagOperation) List() (tags []*models.Tag, err error) {
	err = data.Db.Find(&tags).Error
	return
}

func (_ *tagOperation) Get(tag *models.Tag) (err error) {
	// TODO: how to search using multi query?
	//data.Db.Model(&tag).Related(&users)
	return
}

func (_ *tagOperation) Add(tag *models.Tag) (err error) {
	if err := data.Db.Where("user_id = ? AND value = ?", tag.UserID, tag.Value).
		First(&tag).Error; err != nil {
		return data.Db.Create(tag).Error
	}

	return nil
}

func (_ *tagOperation) remove(id uint) (err error) {
	return data.Db.Where("id = ?", id).Delete(&models.Tag{}).Error
}
