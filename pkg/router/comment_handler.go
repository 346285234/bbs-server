package router

import (
	"encoding/json"
	"github.com/346285234/bbs-server/pkg/bbs"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type CommentResponse struct {
	ID       uint `json:"id"`
	AuthorID uint `json:"author_id"`
	//AuthorName string `json:"author_name"`
	//AuthorPortrait string `json:"author_portrait"`
	Content    string            `json:"content"`
	ModifyTime time.Time         `json:"modify_time"`
	LikeCount  uint              `json:"like_count"`
	Subs       []CommentResponse `json:"sub_comments"`
}

func newCommentResponse(comment bbs.Comment) CommentResponse {
	subComments := make([]CommentResponse, len(comment.Subs))
	for i, v := range comment.Subs {
		subComments[i] = CommentResponse{v.ID, v.AuthorID,
			v.Content, v.UpdatedAt, v.LikeCount,
			[]CommentResponse{}}
	}

	return CommentResponse{comment.ID, comment.AuthorID,
		comment.Content, comment.UpdatedAt, comment.LikeCount,
		subComments}
}

type CommentHandler struct {
	service bbs.CommentService
}

func NewCommentHandler(s bbs.CommentService) CommentHandler {
	return CommentHandler{s}
}

func (c *CommentHandler) List(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	id, _ := strconv.Atoi(p.ByName("topic_id"))
	topicID := uint(id)

	// Get data.
	comments, err := c.service.List(topicID)
	if err != nil {
		return nil, NewAppError(err)
	}

	commentsResponse := make([]CommentResponse, len(comments))
	for i, v := range comments {
		commentsResponse[i] = newCommentResponse(*v)
	}
	data := struct {
		Total  int               `json:"total"`
		Topics []CommentResponse `json:"comments"`
	}{len(commentsResponse), commentsResponse}
	return data, nil
}

func (c *CommentHandler) Reply(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	// request.
	id1, _ := strconv.Atoi(p.ByName("topic_id"))
	id2, _ := strconv.Atoi(r.Header.Get("userID"))
	topicID := uint(id1)
	userID := uint(id2)
	type RequestBody struct {
		ParentID uint `json:"parent_id"`
		Content  string
	}
	var requestBody RequestBody
	json.NewDecoder(r.Body).Decode(&requestBody)
	defer r.Body.Close()

	comment := bbs.Comment{TopicID: topicID, AuthorID: userID, Content: requestBody.Content}

	// db.
	newComment, err := c.service.Reply(comment, requestBody.ParentID)
	if err != nil {
		return nil, NewAppError(err)
	}

	// response.
	data := newCommentResponse(*newComment)
	return data, nil
}

func (_ *CommentHandler) Revoke(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	return nil, nil
}
