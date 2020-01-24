package gorm

import (
	"github.com/346285234/bbs-server/pkg/bbs"
	"github.com/jinzhu/gorm"
)

type LikeService struct {
	op likeOperation
}

func NewLikeService(db *gorm.DB) LikeService {
	op := newLikeOperation(db)
	return LikeService{op}
}

func (l *LikeService) Mark(like bbs.Like, isMark bool) (err error) {
	if isMark {
		err = l.op.add(like)
	} else {
		err = l.op.remove(like)
	}
	// Update like count.
	if like.ObjectType == bbs.TopicType {
		top := newTopicOperation(l.op.db)
		topic, _ := top.get(like.ObjectID)
		if isMark {
			topic.FavoritesCount++
		} else {
			topic.FavoritesCount--
		}
		top.update(topic)
	} else {
		cop := newCommentOperation(l.op.db)
		comment, _ := cop.get(like.ObjectID)
		if isMark {
			comment.LikeCount++
		} else {
			comment.LikeCount--
		}
		cop.update(comment)
	}

	return
}

func (l *LikeService) Check(like bbs.Like) (bool, error) {
	data, err := l.op.get(like)
	if data == nil {
		return false, err
	}
	return true, nil
}

type likeOperation struct {
	db *gorm.DB
}

func newLikeOperation(db *gorm.DB) likeOperation {
	return likeOperation{db}
}

func (l *likeOperation) list() (like []bbs.Like, err error) {
	if err := l.db.Find(&like).Error; err != nil {
		return nil, err
	}

	return like, nil
}

func (l *likeOperation) get(like bbs.Like) (*bbs.Like, error) {
	var result bbs.Like
	if err := l.db.Where("object_type = ? AND object_id = ? AND user_id = ?",
		like.ObjectType, like.ObjectID, like.UserID).
		First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (l *likeOperation) add(like bbs.Like) (err error) {
	return l.db.Create(&like).Error
}

func (l *likeOperation) remove(like bbs.Like) (err error) {
	return l.db.Where("object_type = ? AND object_id = ? AND user_id = ?", like.ObjectType, like.ObjectID, like.UserID).
		Delete(&bbs.Like{}).Error
}
