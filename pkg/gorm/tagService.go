package gorm

import (
	"github.com/346285234/bbs-server/pkg/bbs"
	"github.com/jinzhu/gorm"
)

type TagService struct {
	op tagOperation
}

func NewTagService(db *gorm.DB) TagService {
	op := newTagOperation(db)
	return TagService{op}
}

func (t *TagService) Tags() (categories []*bbs.Tag, err error) {
	return t.op.list()
}

type tagOperation struct {
	db *gorm.DB
}

func newTagOperation(db *gorm.DB) tagOperation {
	return tagOperation{db}
}

func (t *tagOperation) list() (tags []*bbs.Tag, err error) {
	err = t.db.Find(&tags).Error
	return
}

func (t *tagOperation) add(tag *bbs.Tag) (err error) {
	if err := t.db.Where("user_id = ? AND value = ?", tag.UserID, tag.Value).
		First(&tag).Error; err != nil {
		return t.db.Create(tag).Error
	}

	return nil
}

func (t *tagOperation) remove(id uint) (err error) {
	return t.db.Where("id = ?", id).Delete(&bbs.Tag{}).Error
}
