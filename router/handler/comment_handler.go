package handler

import (
	"encoding/json"
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type commentHandler struct {
}

var Ch = commentHandler{}

func (_ *commentHandler) List(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{},*models.AppError) {
	topicID := uint(common.StrToInt(p.ByName("topic_id")))

	// Get data.
	comments, err := services.Cos.List(topicID)
	if err != nil {
		return nil, models.NewAppError(err)
	}

	data := struct {
		Total  int            `json:"total"`
		Topics []*models.Comment `json:"comments"`
	}{len(comments), comments}
	return data, nil
}

func (_ *commentHandler) Reply(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{},*models.AppError) {
	topicID := uint(common.StrToInt(p.ByName("topic_id")))
	userID := uint(common.StrToInt(r.Header.Get("userID")))
	type RequestBody struct {
		ParentID uint `json:"parent_id"`
		Content string
	}
	var requestBody RequestBody
	json.NewDecoder(r.Body).Decode(&requestBody)
	defer r.Body.Close()

	comment := models.Comment{TopicID: topicID, AuthorID:userID, Content:requestBody.Content}
	newComment, err := services.Cos.Reply(comment, requestBody.ParentID)
	if err != nil {
		return nil, models.NewAppError(err)
	}

	return newComment, nil
}

func (_ *commentHandler) Revoke(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{},*models.AppError) {
	return nil, nil
}