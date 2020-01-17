package services

import (
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/operations"
)

type likeService struct {
}

var Ls = likeService{}

func (_ *likeService) Mark(like models.Like, isMark bool) (err error) {
	if isMark {
		err = operations.Lo.Add(like)
	} else {
		err = operations.Lo.Remove(like)
	}
	// Update like count.
	if like.ObjectType == models.TopicType {
		topic, _ := operations.To.Get(like.ObjectID)
		if isMark {
			topic.FavoritesCount++
		} else {
			topic.FavoritesCount--
		}
		operations.To.Update(topic)
	} else {
		comment, _ := operations.CoO.Get(like.ObjectID)
		if isMark {
			comment.LikeCount++
		} else {
			comment.LikeCount--
		}
		operations.CoO.Update(comment)
	}

	return
}

func (_ *likeService) Check(like models.Like) (bool, error) {
	data, err := operations.Lo.Get(like)
	if data == nil {
		return false, err
	}
	return true, nil
}
