package handler

import (
	"encoding/json"
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type tagHandler struct {
}

var TaH = tagHandler{}

func (_ *tagHandler) ListTag(w http.ResponseWriter, r *http.Request, p httprouter.Params) *models.AppError {
	// Get data.
	tags, err := services.TagS.Tags()

	if err != nil {
		return models.NewAppError(err)
	}

	// Set response.
	var response models.Response
	response = models.Response{Success: true, Code: 200, Message: "OK"}
	data := struct {
		Total  int
		Tags []models.Tag
	}{len(tags), tags}
	response.Data = data

	bytes, err := json.Marshal(response)

	if err != nil {
		return models.NewAppError(err)
	}

	w.Write(bytes)

	return nil
}