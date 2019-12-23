package data

type TopicOperation struct {
}

var to = TopicOperation{}

func (_ *TopicOperation) topics() (topics []Topic, err error) {
	db.Find(&topics)
	return topics, nil
}

func (_ *TopicOperation) get(id uint) (topic Topic, err error) {
	db.First(&topic, id)
	return topic, nil
}

func (_ *TopicOperation) add(topic Topic) (err error) {
	db.Create(&topic)
	return nil
}