package handler

import (
	"net/http"

	"github.com/346285234/bbs-server/pkg/router"

	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/pkg/services"
	"github.com/julienschmidt/httprouter"
)

type tagHandler struct {
}

var TaH = tagHandler{}

func (_ *tagHandler) ListTag(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *router.AppError) {
	// db.
	tags, err := services.TagS.Tags()
	if err != nil {
		return nil, router.NewAppError(err)
	}

	// response.
	strings := common.TagsToStrings(tags)
	data := struct {
		Total int      `json:"total"`
		Tags  []string `json:"tags"`
	}{len(strings), strings}
	return data, nil
}
