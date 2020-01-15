package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type EditType int
const (
	Markdown EditType = 1
)

var Models = []interface{}{&Topic{}, &Category{}, &Tag{}, &TopicFavorite{}, &TopicLike{}, &Comment{}}

type Topic struct {
	gorm.Model
	UserID uint `json:"author_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Intro string `json:"description"`
	GroupID uint `json:"group_id"`
	IsPaste bool `json:"is_paste"`
	EditTime time.Duration `json:"edit_time"`
	EditTypeValue int `json:"edit_type"`
	//Comments []Comment // has many
	Tags       []Tag    `gorm:"many2many:topic_tags;"` // many to many
	Category   Category // belong to
	CategoryID uint     `json:"category_id"`
	FavouritesCount uint `json:"favorite_count"`
	LikeCount uint `json:"like_count"`
	ClickCount uint `json:"view_count"`
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
	Topics []Topic `gorm:"many2many:topic_tags;"`
}

type Category struct {
	gorm.Model
	Value string
	topics []Topic
}

type Comment struct {
	gorm.Model
	TopicID uint
	AuthorID uint
	Content string
	LikeCount uint
	Subs []*Comment `gorm:"many2many:subcomments;association_jointable_foreignkey:sub_id"`
}

//type CommentLike struct {
//	gorm.Model
//	CommentID uint
//	UserID uint
//}