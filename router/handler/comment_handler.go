package handler

import (
	"github.com/346285234/bbs-server/data/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type commentHandler struct {
}

var Ch = commentHandler{}

func (_ *commentHandler) List(w http.ResponseWriter, r *http.Request, p httprouter.Params) *models.AppError {
	//topicID := uint(common.StrToInt(p.ByName("topic_id")))

	// Get data.
	//comments, err := services.Cos.List(topicID)
	//if err != nil {
	//	return models.NewAppError(err)
	//}


	return nil
}

func (_ *commentHandler) Reply(w http.ResponseWriter, r *http.Request, p httprouter.Params) *models.AppError {
	return nil
}

func (_ *commentHandler) Revoke(w http.ResponseWriter, r *http.Request, p httprouter.Params) *models.AppError {
	return nil
}