package data

type TopicService struct {
}

var Ts = TopicService{}

func (_ *TopicService) Topics() (topics []Topic, err error) {
	return to.topics()
}

func (_ *TopicService) GetTopic(id uint) (topic *Topic, err error) {
	return to.get(id)
}

func (_ *TopicService) AddTopic(topic Topic) (err error) {
	return to.add(topic)
}

func (_ *TopicService) RemoveTopic(userID uint, topicID uint) (err error) {
	return to.remove(userID, topicID)
}