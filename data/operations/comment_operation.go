package operations

import (
	"github.com/346285234/bbs-server/data/models"
)

type commentOperation struct {
}

var CoO = commentOperation{}


func (_ *commentOperation) List(topicID uint) (topics []models.Comment, err error) {
	//if err := data.Db.Find(&topics).Error; err != nil {
	//	return nil, err
	//}

	return topics, nil
}

func (_ *commentOperation) Add(topicID uint, parentID uint, content string) (err error) {

	return nil
}

func (_ *commentOperation) Remove(topicID uint, id uint) (err error) {
	//if err := data.Db.Where("user_id = ? AND id = ?", userID, topicID).Delete(&models.Topic{}).Error; err != nil {
	//	return err
	//}

	return nil
}