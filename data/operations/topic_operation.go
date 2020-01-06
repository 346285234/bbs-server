package operations

import "github.com/346285234/bbs-server/data"

type TopicOperation struct {
}

var To = TopicOperation{}

func (_ *TopicOperation) Topics() (topics []data.Topic, err error) {
	if err := data.Db.Find(&topics).Error; err != nil {
		return nil, err
	}

	return topics, nil
}

func (_ *TopicOperation) Get(id uint) (topic *data.Topic, err error) {
	var temp data.Topic
	if err := data.Db.First(&temp, id).Error; err != nil {
		return nil, err
	}

	return &temp, nil
}

func (_ *TopicOperation) Add(topic data.Topic) (err error) {
	if err := data.Db.Create(&topic).Error; err != nil {
		return err
	}

	return nil
}

func (_ *TopicOperation) Remove(userID uint, topicID uint) (err error) {
	if err := data.Db.Where("user_id = ? AND id = ?", userID, topicID).Delete(&data.Topic{}).Error; err != nil {
		return err
	}

	return nil
}

// TODO: update topic.
func (_ *TopicOperation) Update() (err error) {

	return nil
}
