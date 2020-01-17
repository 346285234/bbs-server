package services

import (
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/operations"
)

type TopicService struct {
}

var Ts = TopicService{}

func (_ *TopicService) Topics(query map[string]interface{}) (topics []models.Topic, err error) {
	return operations.To.List(query)
}

func (_ *TopicService) GetTopic(id uint) (topic *models.Topic, err error) {
	topic, err = operations.To.Get(id)
	topic.ViewCount++
	operations.To.Update(topic)
	return
}

func (_ *TopicService) AddTopic(topic *models.Topic) error {
	// Add tag if not exist.
	for _, tag := range topic.Tags {
		operations.TagO.Add(tag)
	}

	return operations.To.Add(topic)
}

func (_ *TopicService) RemoveTopic(userID uint, topicID uint) (err error) {
	topic, err := operations.To.Get(topicID)
	if topic.UserID != userID {
		// No permission.
		return
	}

	return operations.To.Remove(topic)
}

func (_ *TopicService) UpdateTopic(topic models.Topic) (*models.Topic, error) {
	// Add tag if not exist.
	for _, tag := range topic.Tags {
		operations.TagO.Add(tag)
	}

	old, err := operations.To.Get(topic.ID)
	old.Title = topic.Title
	old.Content = topic.Content
	old.CategoryID = topic.CategoryID
	old.EditTime = topic.EditTime
	old.EditType = topic.EditType
	old.IsPaste = topic.IsPaste
	old.GroupID = topic.GroupID
	old.Tags = topic.Tags

	err = operations.To.Update(old)
	return old, err
}
