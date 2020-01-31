package router

import (
	"encoding/json"
	"github.com/346285234/bbs-server/pkg/bbs"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type LikeHandler struct {
	service bbs.LikeService
}

func NewLikeHandler(s bbs.LikeService) LikeHandler {
	return LikeHandler{s}
}

func (l *LikeHandler) MarkLikeTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	// request.
	id1, _ := strconv.Atoi(p.ByName("topic_id"))
	id2, _ := strconv.Atoi(r.Header.Get("userID"))
	topicID := uint(id1)
	userID := uint(id2)
	var like bbs.Like
	like.ObjectType = bbs.TopicType
	like.ObjectID = topicID
	like.UserID = userID

	type RequestBody struct {
		IsMark bool `json:"is_mark"`
	}
	var body RequestBody
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	// db.
	err := l.service.Mark(like, body.IsMark)
	if err != nil {
		return nil, NewAppError(err)
	}

	return nil, nil
}

func (l *LikeHandler) CheckLikeTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	id1, _ := strconv.Atoi(p.ByName("topic_id"))
	id2, _ := strconv.Atoi(r.Header.Get("userID"))
	topicID := uint(id1)
	userID := uint(id2)

	var like bbs.Like
	like.ObjectType = bbs.TopicType
	like.ObjectID = topicID
	like.UserID = userID

	isMark, _ := l.service.Check(like)
	var data = struct {
		IsMark bool `json:"is_mark"`
	}{isMark}

	return data, nil
}

func (l *LikeHandler) MarkLikeComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {

	id1, _ := strconv.Atoi(p.ByName("comment_id"))
	id2, _ := strconv.Atoi(r.Header.Get("userID"))
	commentID := uint(id1)
	userID := uint(id2)
	var like bbs.Like
	like.ObjectType = bbs.CommentType
	like.ObjectID = commentID
	like.UserID = userID

	type RequestBody struct {
		IsMark bool `json:"is_mark"`
	}
	var body RequestBody
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	// Get data.
	err := l.service.Mark(like, body.IsMark)

	if err != nil {
		return nil, NewAppError(err)
	}

	return nil, nil
}

func (l *LikeHandler) CheckLikeComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	id1, _ := strconv.Atoi(p.ByName("topic_id"))
	id2, _ := strconv.Atoi(r.Header.Get("userID"))
	commentID := uint(id1)
	userID := uint(id2)

	var like bbs.Like
	like.ObjectType = bbs.CommentType
	like.ObjectID = commentID
	like.UserID = userID

	isMark, _ := l.service.Check(like)
	var data = struct {
		IsMark bool `json:"is_mark"`
	}{isMark}

	return data, nil
}

func (l *LikeHandler) likeTopicUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	id1, _ := strconv.Atoi(p.ByName("topic_id"))
	topicID := uint(id1)

	likes, _ := l.service.List(bbs.TopicType, topicID)
	// TODO: Get users.
	var users = make([]User, len(likes))
	for i, v := range likes {
		users[i] = User{ID: v.UserID}
	}

	var data = struct {
		Users []User `json:"users"`
	}{users}

	return data, nil
}

func (l *LikeHandler) likeCommentUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	id1, _ := strconv.Atoi(p.ByName("comment_id"))
	commentID := uint(id1)

	comments, _ := l.service.List(bbs.CommentType, commentID)
	// TODO: Get users.
	var users = make([]User, len(comments))
	for i, v := range comments {
		users[i] = User{ID: v.UserID}
	}

	var data = struct {
		Users []User `json:"users"`
	}{users}

	return data, nil
}
