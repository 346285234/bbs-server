package handler

import (
	"encoding/json"
	"net/http"

	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/pkg/models"
	"github.com/346285234/bbs-server/pkg/router"
	"github.com/346285234/bbs-server/pkg/services"
	"github.com/julienschmidt/httprouter"
)

type commentHandler struct {
}

var Ch = commentHandler{}

func (_ *commentHandler) List(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *router.AppError) {
	topicID := uint(common.StrToInt(p.ByName("topic_id")))

	// Get data.
	comments, err := services.Cos.List(topicID)
	if err != nil {
		return nil, router.NewAppError(err)
	}

	commentsResponse := make([]router.CommentResponse, len(comments))
	for i, v := range comments {
		commentsResponse[i] = common.CommentToResponse(*v)
	}
	data := struct {
		Total  int                      `json:"total"`
		Topics []router.CommentResponse `json:"comments"`
	}{len(commentsResponse), commentsResponse}
	return data, nil
}

func (_ *commentHandler) Reply(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *router.AppError) {
	// request.
	topicID := uint(common.StrToInt(p.ByName("topic_id")))
	userID := uint(common.StrToInt(r.Header.Get("userID")))
	type RequestBody struct {
		ParentID uint `json:"parent_id"`
		Content  string
	}
	var requestBody RequestBody
	json.NewDecoder(r.Body).Decode(&requestBody)
	defer r.Body.Close()

	comment := models.Comment{TopicID: topicID, AuthorID: userID, Content: requestBody.Content}

	// db.
	newComment, err := services.Cos.Reply(comment, requestBody.ParentID)
	if err != nil {
		return nil, router.NewAppError(err)
	}

	// response.
	data := common.CommentToResponse(*newComment)
	return data, nil
}

func (_ *commentHandler) Revoke(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *router.AppError) {
	return nil, nil
}
