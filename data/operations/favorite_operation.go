package operations

import (
	"github.com/346285234/bbs-server/data"
	"github.com/346285234/bbs-server/data/models"
)

type favoriteOperation struct {
}

var Fo = &favoriteOperation{}

func (_ *favoriteOperation) List() (favorites []models.TopicFavorite, err error) {
	if err := data.Db.Find(&favorites).Error; err != nil {
		return nil, err
	}

	return favorites, nil
}

func (_ *favoriteOperation) Get(userID uint, topicID uint) (favorite *models.TopicFavorite, err error) {
	var result models.TopicFavorite
	if err := data.Db.Where("user_id = ? AND topic_id = ?", userID, topicID).
		First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (_ *favoriteOperation) Add(favorite models.TopicFavorite) (err error) {
	return data.Db.Create(&favorite).Error
}

func (_ *favoriteOperation) Remove(favorite models.TopicFavorite) (err error) {
	return data.Db.Where("topic_id = ? AND user_id = ?", favorite.TopicID, favorite.UserID).
		Delete(&models.TopicFavorite{}).Error
}