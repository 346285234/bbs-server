package handler

import (
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type tagHandler struct {
}

var TaH = tagHandler{}

func (_ *tagHandler) ListTag(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {
	// db.
	tags, err := services.TagS.Tags()
	if err != nil {
		return nil, models.NewAppError(err)
	}

	// response.
	strings := common.TagsToStrings(tags)
	data := struct {
		Total int      `json:"total"`
		Tags  []string `json:"tags"`
	}{len(strings), strings}
	return data, nil
}
