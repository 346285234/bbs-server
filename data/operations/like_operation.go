package operations

import (
	"github.com/346285234/bbs-server/data"
	"github.com/346285234/bbs-server/data/models"
)

type likeOperation struct {
}

var Lo = &likeOperation{}

func (_ *likeOperation) List() (like []models.Like, err error) {
	if err := data.Db.Find(&like).Error; err != nil {
		return nil, err
	}

	return like, nil
}

func (_ *likeOperation) Get(like models.Like) (*models.Like, error) {
	var result models.Like
	if err := data.Db.Where("object_type = ? AND object_id = ? AND user_id = ?",
		like.ObjectType, like.ObjectID, like.UserID).
		First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (_ *likeOperation) Add(like models.Like) (err error) {
	return data.Db.Create(&like).Error
}

func (_ *likeOperation) Remove(like models.Like) (err error) {
	return data.Db.Where("object_type = ? AND object_id = ? AND user_id = ?", like.ObjectType, like.ObjectID, like.UserID).
		Delete(&models.Like{}).Error
}
