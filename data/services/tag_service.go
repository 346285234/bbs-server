package services

import (
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/operations"
)

type tagService struct {
}

var TagS = tagService{}

func (_ *tagService) Tags() (categories []*models.Tag, err error) {
	return operations.TagO.List()
}
