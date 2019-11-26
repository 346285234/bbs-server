package data

import (
	"github.com/jinzhu/gorm"
)

type Topic struct {
	gorm.Model
	AuthorID uint
	title string
	Intro string
	Content string

	Comments []Comment

	Tags []*Tag `gorm:"many2many:topic_tags;"`
	CategoryID uint
	FavoriteCount uint
	LikeCount uint
	ClickCount uint
}

type TopicFavorite struct {
	gorm.Model
	TopicID uint
	UserID uint
}

type TopicLike struct {
	gorm.Model
	TopicID uint
	UserID uint
}

type Tag struct {
	gorm.Model
	UserID uint
	Value string
	Topics []*Topic `gorm:"many2many:topic_tags;"`
}

type Category struct {
	gorm.Model
	Value string
	topics []Topic
}

type Comment struct {
	gorm.Model
	AuthorID uint
	TopicID uint
	Content string
	LikeCount uint
	Children []Comment `gorm:"foreignkey:Parent"`
	Parent uint
}

type CommentLike struct {
	gorm.Model
	CommentID uint
	UserID uint
}