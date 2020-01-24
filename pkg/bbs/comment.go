package bbs

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	TopicID   uint
	AuthorID  uint
	Content   string
	LikeCount uint
	Subs      []*Comment `gorm:"many2many:subcomments;association_jointable_foreignkey:sub_id"`
}

type CommentService interface {
	List(topicID uint) (comments []*Comment, err error)
	Reply(comment Comment, parentID uint) (*Comment, error)
	Revoke(topicID uint, id uint) (err error)
}
