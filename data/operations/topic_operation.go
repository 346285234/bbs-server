package operations

import (
	"github.com/346285234/bbs-server/data"
	"github.com/346285234/bbs-server/data/models"
)

type topicOperation struct {
}

var To = newTopicOperation()

func newTopicOperation() *topicOperation {
	return &topicOperation{}
}

func (_ *topicOperation) List() (topics []models.Topic, err error) {
	if err := data.Db.Find(&topics).Error; err != nil {
		return nil, err
	}

	return topics, nil
}

func (_ *topicOperation) Get(id uint) (topic *models.Topic, err error) {
	var temp models.Topic
	if err := data.Db.First(&temp, id).Error; err != nil {
		return nil, err
	}

	return &temp, nil
}

func (_ *topicOperation) Add(topic models.Topic) (err error) {
	if err := data.Db.Create(&topic).Error; err != nil {
		return err
	}

	return nil
}

func (_ *topicOperation) Remove(userID uint, topicID uint) (err error) {
	if err := data.Db.Where("user_id = ? AND id = ?", userID, topicID).Delete(&models.Topic{}).Error; err != nil {
		return err
	}

	return nil
}

// TODO: update topic.
func (_ *topicOperation) Update() (err error) {

	return nil
}
