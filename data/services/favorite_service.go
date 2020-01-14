package services

import (
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/operations"
)

type FavoriteService struct {
}

var Fs = FavoriteService{}

func (_ *FavoriteService) Mark(favorite models.TopicFavorite, isMark bool) (err error) {
	if isMark {
		return operations.Fo.Add(favorite)
	} else {
		return operations.Fo.Remove(favorite)
	}
	// Update favorite count.
}

func (_ *FavoriteService) Check(favorite models.TopicFavorite) (err error) {
	_, err = operations.Fo.Get(favorite.UserID, favorite.TopicID)
	return err
}