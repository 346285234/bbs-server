package services

import (
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/operations"
)

type TopicService struct {

}

var Ts = TopicService{}

func (_ *TopicService) Topics() (topics []models.Topic, err error) {
	return operations.To.List()
}

func (_ *TopicService) GetTopic(id uint) (topic *models.Topic, err error) {
	return operations.To.Get(id)
}

func (_ *TopicService) AddTopic(topic models.Topic) (err error) {
	return operations.To.Add(topic)
}

func (_ *TopicService) RemoveTopic(userID uint, topicID uint) (err error) {
	return operations.To.Remove(userID, topicID)
}
