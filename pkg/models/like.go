package models

import "github.com/jinzhu/gorm"

type ObjectType uint

const (
	TopicType = iota
	CommentType
)

type Like struct {
	gorm.Model
	ObjectType ObjectType
	ObjectID   uint
	UserID     uint
}

type LikeService interface {
	Mark(like Like, isMark bool) (err error)
	Check(like Like) (bool, error)
}
