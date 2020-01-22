package gorm

import (
	"github.com/346285234/bbs-server/pkg/models"
	"github.com/jinzhu/gorm"
)

type TopicService struct {
	op topicOperation
}

func NewTopicService(db *gorm.DB) TopicService {
	op := newTopicOperation(db)
	return TopicService{op}
}

func (t *TopicService) Topics(query map[string]interface{}) (topics []models.Topic, err error) {
	return t.op.list(query)
}

func (t *TopicService) GetTopic(id uint) (topic *models.Topic, err error) {
	topic, err = t.op.get(id)
	topic.ViewCount++
	t.op.update(topic)
	return
}

func (t *TopicService) AddTopic(topic *models.Topic) error {
	// Add tag if not exist.
	top := newTagOperation(t.op.db)
	for _, tag := range topic.Tags {
		top.add(tag)
	}

	return t.op.add(topic)
}

func (t *TopicService) RemoveTopic(userID uint, topicID uint) (err error) {
	topic, err := t.op.get(topicID)
	if topic.UserID != userID {
		// No permission.
		return
	}

	return t.op.remove(topic)
}

func (t *TopicService) UpdateTopic(topic models.Topic) (*models.Topic, error) {
	// Add tag if not exist.
	top := newTagOperation(t.op.db)
	for _, tag := range topic.Tags {
		top.add(tag)
	}

	old, err := t.op.get(topic.ID)
	old.Title = topic.Title
	old.Content = topic.Content
	old.CategoryID = topic.CategoryID
	old.EditTime = topic.EditTime
	old.EditType = topic.EditType
	old.IsPaste = topic.IsPaste
	old.GroupID = topic.GroupID
	old.Tags = topic.Tags

	err = t.op.update(old)
	return old, err
}

type topicOperation struct {
	db *gorm.DB
}

func newTopicOperation(db *gorm.DB) topicOperation {
	return topicOperation{db}
}

// List topics using query.
func (t *topicOperation) list(query map[string]interface{}) (topics []models.Topic, err error) {
	var db = t.db
	if v, ok := query["tag"]; ok {
		delete(query, "tag")
		db = db.Joins("join topic_tags on topics.id = topic_tags.topic_id "+
			"join tags on topic_tags.tag_id = tags.id").
			Where("tags.value = ?", v)
	}

	var page, pageSize uint
	if _, ok := query["page"]; ok {
		page, _ = query["page"].(uint)
		pageSize, _ = query["page_size"].(uint)
		delete(query, "page")
		delete(query, "page_size")
	}

	db = db.Where(query)

	if page != 0 && pageSize != 0 {
		db = db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	if err := db.Preload("Tags").
		Preload("Category").
		Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}

func (t *topicOperation) get(id uint) (*models.Topic, error) {
	var result models.Topic
	if err := t.db.Preload("Tags").Preload("Category").First(&result, id).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (t *topicOperation) add(topic *models.Topic) error {
	var category models.Category
	t.db.Model(&topic).Related(&category)
	topic.Category = category
	if err := t.db.Create(&topic).Error; err != nil {
		return err
	}
	return nil
}

func (t *topicOperation) remove(topic *models.Topic) (err error) {
	t.db.Model(topic).Association("Tags").Clear()
	t.db.Delete(topic)
	return
}

func (t *topicOperation) update(topic *models.Topic) error {
	var category models.Category
	t.db.Model(topic).Related(&category)
	topic.Category = category
	if err := t.db.Save(topic).Error; err != nil {
		return err
	}
	t.db.Model(topic).Association("Tags").Replace(topic.Tags)

	return nil
}
