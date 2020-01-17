package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

var Models = []interface{}{&Topic{}, &Category{}, &Tag{}, &Favorite{}, &Like{}, &Comment{}}

type ObjectType uint

const (
	TopicType = iota
	CommentType
)

type Topic struct {
	gorm.Model
	UserID         uint
	Title          string
	Content        string
	Intro          string
	GroupID        uint
	IsPaste        bool
	EditTime       time.Duration
	EditType       uint
	Comments       []Comment // has many
	Tags           []*Tag    `gorm:"many2many:topic_tags;"` // many to many
	Category       Category  // belong to
	CategoryID     uint
	FavoritesCount uint
	LikeCount      uint
	ViewCount      uint
}

type Favorite struct {
	gorm.Model
	TopicID uint
	UserID  uint
}

type Like struct {
	gorm.Model
	ObjectType ObjectType
	ObjectID   uint
	UserID     uint
}

type Tag struct {
	gorm.Model
	UserID uint
	Value  string
	Topics []*Topic `gorm:"many2many:topic_tags;"`
}

type Category struct {
	gorm.Model
	Value  string
	topics []Topic
}

type Comment struct {
	gorm.Model
	TopicID   uint
	AuthorID  uint
	Content   string
	LikeCount uint
	Subs      []*Comment `gorm:"many2many:subcomments;association_jointable_foreignkey:sub_id"`
}
