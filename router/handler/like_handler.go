package handler

import (
	"encoding/json"
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type likeHandler struct {
}

var LiH = likeHandler{}

func (_ *likeHandler) MarkLike(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{},*models.AppError) {

	topicID := uint(common.StrToInt(p.ByName("topic_id")))
	userID := uint(common.StrToInt(r.Header.Get("userID")))
	var like models.TopicLike
	like.TopicID = topicID
	like.UserID = userID

	type RequestBody struct {
		Unmark bool `json:"unmark"`
	}
	var body RequestBody
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	// Get data.
	err := services.Ls.Mark(like, !body.Unmark)

	if err != nil {
		return nil, models.NewAppError(err)
	}

	return nil, nil
}

func (_ *likeHandler) CheckLike(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{},*models.AppError) {
	topicID := uint(common.StrToInt(p.ByName("topic_id")))
	userID := uint(common.StrToInt(r.Header.Get("userID")))

	var like models.TopicLike
	like.TopicID = topicID
	like.UserID = userID

	err := services.Ls.Check(like)
	var isMark bool
	if err == nil {
		isMark = true
	} else {
		isMark = false
	}
	var data = struct{
		IsMark bool `json:"is_mark"`
	}{isMark}

	return data, nil
}