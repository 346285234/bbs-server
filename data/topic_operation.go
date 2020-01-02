package data

type TopicOperation struct {
}

var to = TopicOperation{}

func (_ *TopicOperation) topics() (topics []Topic, err error) {
	if err := db.Find(&topics).Error; err != nil {
		return nil, err
	}

	return topics, nil
}

func (_ *TopicOperation) get(id uint) (topic *Topic, err error) {
	var temp Topic
	if err := db.First(&temp, id).Error; err != nil {
		return nil, err
	}

	return &temp, nil
}

func (_ *TopicOperation) add(topic Topic) (err error) {
	if err := db.Create(&topic).Error; err != nil {
		return err
	}

	return nil
}

func (_ *TopicOperation) remove(userID uint, topicID uint) (err error) {
	if err := db.Where("author_id = ? AND id = ?", userID, topicID).Delete(&Topic{}).Error; err != nil {
		return err
	}

	return nil
}

// TODO: update topic.
func (_ *TopicOperation) update() (err error) {

	return nil
}
