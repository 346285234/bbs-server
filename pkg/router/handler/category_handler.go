// Package handler implements all http handler.
package handler

import (
	"net/http"

	"github.com/346285234/bbs-server/pkg/models"

	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/pkg/router"
	"github.com/346285234/bbs-server/pkg/services"
	"github.com/julienschmidt/httprouter"
)

type categoryHandler struct {
	service models.CategoryService
}

var CaH = categoryHandler{}

func (_ *categoryHandler) ListCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *router.AppError) {
	// db.
	categories, err := services.Cs.Categories()
	if err != nil {
		return nil, router.NewAppError(err)
	}

	// response.
	response := common.CategoriesToResponse(categories)
	data := struct {
		Total      int                       `json:"total"`
		Categories []router.CategoryResponse `json:"categories"`
	}{len(response), response}
	return data, nil
}
