package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type EditType int
const (
	Markdown EditType = 1
)

type Topic struct {
	gorm.Model
	UserID uint
	Title string
	Content string
	Intro string
	GroupID uint `json:"group_id"`
	IsPaste bool `json:"is_paste"`
	EditTime time.Duration `json:"edit_time"`
	EditTypeValue int `json:"edit_type"`
	//Comments []Comment // has many
	Tags       []Tag    `gorm:"many2many:topic_tags;"` // many to many
	Category   Category // belong to
	CategoryID uint     `json:"category_id"`
	//FavouritesCount uint
	//LikeCount uint
	//ClickCount uint
}

//type TopicFavorite struct {
//	gorm.Model
//	TopicID uint
//	UserID uint
//}
//
//type TopicLike struct {
//	gorm.Model
//	TopicID uint
//	UserID uint
//}

type Tag struct {
	gorm.Model
	UserID uint
	Value string
	Topics []Topic `gorm:"many2many:topic_tags;"`
}

type Category struct {
	gorm.Model
	Value string
	topics []Topic
}

//type Comment struct {
//	gorm.Model
//	AuthorID uint
//	Content string
//	LikeCount uint
//	Children []Comment `gorm:"foreignkey:Parent"` // has many
//	Parent uint
//	TopicID uint
//}
//
//type CommentLike struct {
//	gorm.Model
//	CommentID uint
//	UserID uint
//}