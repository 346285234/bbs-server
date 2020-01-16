package handler

import (
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type categoryHandler struct {
}

var CaH = categoryHandler{}

func (_ *categoryHandler) ListCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {
	// db.
	categories, err := services.Cs.Categories()
	if err != nil {
		return nil, models.NewAppError(err)
	}

	// response.
	response := common.CategoriesToResponse(categories)
	data := struct {
		Total      int                       `json:"total"`
		Categories []models.CategoryResponse `json:"categories"`
	}{len(response), response}
	return data, nil
}
