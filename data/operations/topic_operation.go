package operations

import "github.com/346285234/bbs-server/data"

type TopicOperation struct {
}

var To = TopicOperation{}

func (_ *TopicOperation) Topics() (topics *[]data.Topic, err error) {
	return nil, nil
}

func (_ *TopicOperation) Get(id uint) (topic *data.Topic, err error) {
	return nil, nil
}
