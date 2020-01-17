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

func (_ *topicOperation) List(query map[string]interface{}) (topics []models.Topic, err error) {
	if err := data.Db.Preload("Tags").
		Preload("Category").Where(query).
		Find(&topics).Error; err != nil {
		return nil, err
	}

	return topics, nil
}

func (_ *topicOperation) Get(id uint) (*models.Topic, error) {
	var result models.Topic
	if err := data.Db.Preload("Tags").Preload("Category").First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (_ *topicOperation) Add(topic *models.Topic) error {
	var category models.Category
	data.Db.Model(&topic).Related(&category)
	topic.Category = category
	if err := data.Db.Create(&topic).Error; err != nil {
		return err
	}
	return nil
}

func (_ *topicOperation) Remove(topic *models.Topic) (err error) {
	data.Db.Model(topic).Association("Tags").Clear()
	data.Db.Delete(topic)
	return
}

func (_ *topicOperation) Update(topic *models.Topic) error {
	var category models.Category
	data.Db.Model(topic).Related(&category)
	topic.Category = category
	if err := data.Db.Save(topic).Error; err != nil {
		return err
	}
	data.Db.Model(topic).Association("Tags").Replace(topic.Tags)

	return nil
}
