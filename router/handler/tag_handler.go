package handler

import (
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type tagHandler struct {
}

var TaH = tagHandler{}

func (_ *tagHandler) ListTag(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{},*models.AppError) {
	// Get data.
	tags, err := services.TagS.Tags()

	if err != nil {
		return nil, models.NewAppError(err)
	}

	data := struct {
		Total  int
		Tags []models.Tag
	}{len(tags), tags}

	return data, nil
}