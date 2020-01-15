package services

import (
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/operations"
)

type commentService struct {
}

var Cos = commentService{}

func (_ *commentService) List(topicID uint) (comments []*models.Comment, err error) {
	return operations.CoO.List(topicID)
}

func (_ *commentService) Reply(comment models.Comment, parentID uint) (*models.Comment, error) {
	return operations.CoO.Add(comment, parentID)
}

func (_ *commentService) Revoke(topicID uint, id uint) (err error) {
	return operations.CoO.Remove(topicID, id)
}