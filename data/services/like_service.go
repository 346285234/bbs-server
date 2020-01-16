package services

import (
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/operations"
)

type likeService struct {
}

var Ls = likeService{}

func (_ *likeService) Mark(like models.Like, isMark bool) error {
	if isMark {
		return operations.Lo.Add(like)
	} else {
		return operations.Lo.Remove(like)
	}
	// Update like count.
}

func (_ *likeService) Check(like models.Like) (bool, error) {
	data, err := operations.Lo.Get(like)
	if data == nil {
		return false, err
	}
	return true, nil
}
