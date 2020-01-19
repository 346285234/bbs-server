/*
Package services is used to operate data indirectly.
*/
package services

import (
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/operations"
)

type CategoryService struct {
}

var Cs = CategoryService{}

func (_ *CategoryService) Categories() (categories []models.Category, err error) {
	return operations.Co.List()
}
