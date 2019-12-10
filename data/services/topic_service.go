package services

import (
	"github.com/346285234/bbs-server/data"
	"github.com/346285234/bbs-server/data/operations"
)

type TopicService struct {
}

var Ts = TopicService{}

func (_ *TopicService) Topics() (topics *[]data.Topic, err error) {
	return operations.To.Topics()
}

func (_ *TopicService) GetTopic(id uint) (topic *data.Topic, err error) {
	return operations.To.Get(id)
}