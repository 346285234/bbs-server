package database

import (
	"github.com/346285234/bbs-server/pkg/bbs"
	"github.com/jinzhu/gorm"
)

type FollowService struct {
	op followOperation
}

func NewFollowService(db *gorm.DB) FollowService {
	op := newFollowOperation(db)
	return FollowService{op}
}

func (f *FollowService) Mark(follow bbs.Follow, isMark bool) (err error) {
	// Change favorite table.
	if isMark {
		err = f.op.add(follow)
	} else {
		err = f.op.remove(follow)
	}

	return
}

func (f *FollowService) Check(follow bbs.Follow) (bool, error) {
	data, err := f.op.get(follow)
	if data == nil {
		return false, err
	}
	return true, nil
}

func (f *FollowService) List(userID uint) ([]bbs.Follow, error) {
	return f.op.list(userID)
}

type followOperation struct {
	db *gorm.DB
}

func newFollowOperation(db *gorm.DB) followOperation {
	return followOperation{db}
}

func (f *followOperation) listAll() (follows []bbs.Follow, err error) {
	if err := f.db.Find(&follows).Error; err != nil {
		return nil, err
	}

	return follows, nil
}

func (f *followOperation) list(userID uint) (follows []bbs.Follow, err error) {
	if err := f.db.Where("subject_id = ?", userID).Find(&follows).Error; err != nil {
		return nil, err
	}

	return
}

func (f *followOperation) get(follow bbs.Follow) (*bbs.Follow, error) {
	var result bbs.Follow
	if err := f.db.Where("subject_id = ? AND object_id = ?", follow.SubjectID, follow.ObjectID).
		First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (f *followOperation) add(follow bbs.Follow) (err error) {
	return f.db.Create(&follow).Error
}

func (f *followOperation) remove(follow bbs.Follow) (err error) {
	return f.db.Where("subject_id = ? AND object_id = ?", follow.SubjectID, follow.ObjectID).
		Delete(&bbs.Follow{}).Error
}
