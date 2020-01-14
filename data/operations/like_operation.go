package operations

import (
	"github.com/346285234/bbs-server/data"
	"github.com/346285234/bbs-server/data/models"
)

type likeOperation struct {
}

var Lo = &likeOperation{}

func (_ *likeOperation) List() (like []models.TopicLike, err error) {
	if err := data.Db.Find(&like).Error; err != nil {
		return nil, err
	}

	return like, nil
}

func (_ *likeOperation) Get(userID uint, topicID uint) (like *models.TopicLike, err error) {
	var result models.TopicLike
	if err := data.Db.Where("user_id = ? AND topic_id = ?", userID, topicID).
		First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (_ *likeOperation) Add(like models.TopicLike) (err error) {
	return data.Db.Create(&like).Error
}

func (_ *likeOperation) Remove(like models.TopicLike) (err error) {
	return data.Db.Where("topic_id = ? AND user_id = ?", like.TopicID, like.UserID).
		Delete(&models.TopicLike{}).Error
}