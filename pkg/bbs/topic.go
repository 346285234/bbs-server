package bbs

import (
	"time"

	"github.com/jinzhu/gorm"
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

type TopicService interface {
	Topics(query map[string]interface{}) (topics []Topic, err error)
	GetTopic(id uint) (topic *Topic, err error)
	AddTopic(topic *Topic) error
	RemoveTopic(userID uint, topicID uint) (err error)
	UpdateTopic(topic Topic) (*Topic, error)
}
