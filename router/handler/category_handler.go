package handler

import (
	"encoding/json"
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type categoryHandler struct {
}

var CaH = categoryHandler{}

func (_ *categoryHandler) ListCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params) *models.AppError {
	// Get data.
	categories, err := services.Cs.Categories()

	if err != nil {
		return models.NewAppError(err)
	}

	// Set response.
	type Category struct {
		Id uint `json:"id"`
		Name string `json:"name"`
	}
	type Data struct {
		Total  int `json:"total"`
		Categories []Category `json:"categories"`
	}
	values := make([]Category, len(categories))
	for i, v := range categories {
		value := Category{v.ID, v.Value}
		values[i] = value
	}
	data := Data{len(values), values}
	response := models.Response{true, 200, "OK", data}
	bytes, err := json.Marshal(response)
	if err != nil {
		return models.NewAppError(err)
	}
	w.Write(bytes)

	return nil
}