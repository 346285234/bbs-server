package bbs

import "github.com/jinzhu/gorm"

type Follow struct {
	gorm.Model
	SubjectID uint
	ObjectID  uint
}

type FollowService interface {
	Mark(follow Follow, isMark bool) (err error)
	Check(follow Follow) (bool, error)
	List(userID uint) (follows []Follow, err error)
}
