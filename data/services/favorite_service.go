package services

import (
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/operations"
)

type FavoriteService struct {
}

var Fs = FavoriteService{}

func (_ *FavoriteService) Mark(favorite models.Favorite, isMark bool) (err error) {
	// Change favorite table.
	if isMark {
		err = operations.Fo.Add(favorite)
	} else {
		err = operations.Fo.Remove(favorite)
	}
	// Update topic favorite count.
	topic, err := operations.To.Get(favorite.TopicID)
	if isMark {
		topic.FavoritesCount++
	} else {
		topic.FavoritesCount--
	}
	operations.To.Update(topic)
	return
}

func (_ *FavoriteService) Check(favorite models.Favorite) (bool, error) {
	data, err := operations.Fo.Get(favorite)
	if data == nil {
		return false, err
	}
	return true, nil
}
