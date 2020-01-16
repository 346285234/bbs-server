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

func (_ *likeHandler) MarkLikeTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {
	// request.
	topicID := uint(common.StrToInt(p.ByName("topic_id")))
	userID := uint(common.StrToInt(r.Header.Get("userID")))
	var like models.Like
	like.ObjectType = models.TopicType
	like.ObjectID = topicID
	like.UserID = userID

	type RequestBody struct {
		Unmark bool `json:"unmark"`
	}
	var body RequestBody
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	// db.
	err := services.Ls.Mark(like, !body.Unmark)
	if err != nil {
		return nil, models.NewAppError(err)
	}

	return nil, nil
}

func (_ *likeHandler) CheckLikeTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {
	topicID := uint(common.StrToInt(p.ByName("topic_id")))
	userID := uint(common.StrToInt(r.Header.Get("userID")))

	var like models.Like
	like.ObjectType = models.TopicType
	like.ObjectID = topicID
	like.UserID = userID

	isMark, _ := services.Ls.Check(like)
	var data = struct {
		IsMark bool `json:"is_mark"`
	}{isMark}

	return data, nil
}

func (_ *likeHandler) MarkLikeComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {

	commentID := uint(common.StrToInt(p.ByName("comment_id")))
	userID := uint(common.StrToInt(r.Header.Get("userID")))
	var like models.Like
	like.ObjectType = models.CommentType
	like.ObjectID = commentID
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

func (_ *likeHandler) CheckLikeComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {
	commentID := uint(common.StrToInt(p.ByName("comment_id")))
	userID := uint(common.StrToInt(r.Header.Get("userID")))

	var like models.Like
	like.ObjectType = models.CommentType
	like.ObjectID = commentID
	like.UserID = userID

	isMark, _ := services.Ls.Check(like)
	var data = struct {
		IsMark bool `json:"is_mark"`
	}{isMark}

	return data, nil
}
