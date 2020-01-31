package database

import (
	"github.com/346285234/bbs-server/pkg/bbs"
	"github.com/jinzhu/gorm"
)

type FavoriteService struct {
	op favoriteOperation
}

func NewFavoriteService(db *gorm.DB) FavoriteService {
	op := newFavoriteOperation(db)
	return FavoriteService{op}
}

func (f *FavoriteService) Mark(favorite bbs.Favorite, isMark bool) (err error) {
	// Change favorite table.
	if isMark {
		err = f.op.add(favorite)
	} else {
		err = f.op.remove(favorite)
	}

	// Update topic favorite count.
	top := newTopicOperation(f.op.db)
	topic, err := top.get(favorite.TopicID)
	if isMark {
		topic.FavoritesCount++
	} else {
		topic.FavoritesCount--
	}
	top.update(topic)
	return
}

func (f *FavoriteService) Check(favorite bbs.Favorite) (bool, error) {
	data, err := f.op.get(favorite)
	if data == nil {
		return false, err
	}
	return true, nil
}

func (f *FavoriteService) List(topicID uint) ([]bbs.Favorite, error) {
	return f.op.listWith(topicID)
}

type favoriteOperation struct {
	db *gorm.DB
}

func newFavoriteOperation(db *gorm.DB) favoriteOperation {
	return favoriteOperation{db}
}

func (f *favoriteOperation) list() (favorites []bbs.Favorite, err error) {
	if err := f.db.Find(&favorites).Error; err != nil {
		return nil, err
	}

	return favorites, nil
}

func (f *favoriteOperation) listWith(topicID uint) (favorites []bbs.Favorite, err error) {
	if err := f.db.Where("topic_id = ?", topicID).Find(&favorites).Error; err != nil {
		return nil, err
	}

	return favorites, nil
}

func (f *favoriteOperation) get(favorite bbs.Favorite) (*bbs.Favorite, error) {
	var result bbs.Favorite
	if err := f.db.Where("user_id = ? AND topic_id = ?", favorite.UserID, favorite.TopicID).
		First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (f *favoriteOperation) add(favorite bbs.Favorite) (err error) {
	return f.db.Create(&favorite).Error
}

func (f *favoriteOperation) remove(favorite bbs.Favorite) (err error) {
	return f.db.Where("topic_id = ? AND user_id = ?", favorite.TopicID, favorite.UserID).
		Delete(&bbs.Favorite{}).Error
}
