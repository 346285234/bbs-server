package services

import (
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/operations"
)

type likeService struct {
}

var Ls = likeService{}

func (_ *likeService) Mark(like models.TopicLike, isMark bool) (err error) {
	if isMark {
		return operations.Lo.Add(like)
	} else {
		return operations.Lo.Remove(like)
	}
	// Update topic like count.
}

func (_ *likeService) Check(like models.TopicLike) (err error) {
	_, err = operations.Lo.Get(like.UserID, like.TopicID)
	return err
}