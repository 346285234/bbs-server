package router

import (
	"encoding/json"
	"errors"
	"github.com/346285234/bbs-server/pkg/bbs"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type TopicRequest struct {
	ID         uint
	Title      string
	Content    string
	CategoryID uint `json:"category_id"`
	Tags       tags
	EditTime   time.Duration `json:"edit_time"`
	IsPaste    bool          `json:"is_paste"`
	EditType   uint          `json:"edit_type"`
	GroupID    uint          `json:"group_id"`
}

func (t TopicRequest) RequestToTopic(userID uint) bbs.Topic {
	tags := t.Tags.StringsToTags(userID)
	topic := bbs.Topic{Title: t.Title, Content: t.Content,
		CategoryID: t.CategoryID, Tags: tags,
		EditTime: t.EditTime, IsPaste: t.IsPaste,
		EditType: t.EditType, GroupID: t.GroupID,
		UserID: userID,
	}
	if t.ID != 0 {
		topic.ID = t.ID
	}

	return topic
}

type TopicResponse struct {
	ID           uint     `json:"id"`
	Title        string   `json:"title"`
	Tags         []string `json:"tags"`
	CategoryID   uint     `json:"category_id"`
	CategoryName string   `json:"category_name"`
	AuthorID     uint     `json:"author_id"`
	//AuthorName string `json:"author_name"`
	//AuthorPortrait string `json:"author_portrait"`
	Description   string    `json:"description"`
	Content       string    `json:"content"`
	ModifyTime    time.Time `json:"modify_time"`
	FavoriteCount uint      `json:"favorite_count"`
	LikeCount     uint      `json:"like_count"`
	ViewCount     uint      `json:"view_count"`
}

func Intro(from string) string {
	return from[:1]
}

func newTopicResponse(topic bbs.Topic) TopicResponse {
	strings := newTags(topic.Tags)
	response := TopicResponse{topic.ID, topic.Title, strings,
		topic.CategoryID, topic.Category.Value,
		topic.UserID, topic.Intro, topic.Content,
		topic.UpdatedAt, topic.FavoritesCount,
		topic.LikeCount, topic.ViewCount}

	return response
}

type TopicHandler struct {
	service bbs.TopicService
}

func NewTopicHandler(s bbs.TopicService) TopicHandler {
	return TopicHandler{s}
}

func (t *TopicHandler) ListTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	// TODO: request.

	vars := r.URL.Query()

	query := make(map[string]interface{})
	id1, _ := strconv.Atoi(p.ByName("topic_id"))
	groupID := uint(id1)
	if groupID != 0 {
		query["group_id"] = groupID
	}
	id2, _ := strconv.Atoi(p.ByName("user_id"))
	userID := uint(id2)
	if userID != 0 {
		query["user_id"] = userID
	}
	id3, _ := strconv.Atoi(p.ByName("category_id"))
	categoryID := uint(id3)
	if categoryID != 0 {
		query["category_id"] = categoryID
	}
	id4, _ := strconv.Atoi(p.ByName("page"))
	id5, _ := strconv.Atoi(p.ByName("page_size"))
	page := uint(id4)
	pageSize := uint(id5)
	if page != 0 && pageSize != 0 {
		query["page"] = page
		query["page_size"] = pageSize
	}

	tag := vars.Get("tag")
	if tag != "" {
		query["tag"] = tag
	}

	// db.
	topics, err := t.service.Topics(query)
	if err != nil {
		return nil, NewAppError(err)
	}

	// response.
	topicsResponse := make([]TopicResponse, len(topics))
	for i, v := range topics {
		topicResponse := newTopicResponse(v)
		topicsResponse[i] = topicResponse
	}

	data := struct {
		Total  int             `json:"total"`
		Topics []TopicResponse `json:"topics"`
	}{len(topicsResponse), topicsResponse}
	return data, nil
}

func (t *TopicHandler) GetTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {

	// request.
	id1, _ := strconv.Atoi(p.ByName("topic_id"))
	id := uint(id1)

	// db.
	topic, err := t.service.GetTopic(id)
	if err != nil {
		return nil, NewAppError(err)
	}

	// response.
	topicResponse := newTopicResponse(*topic)
	return topicResponse, nil
}

func (t *TopicHandler) AddTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	// request.
	id1, _ := strconv.Atoi(p.ByName("userID"))
	userID := uint(id1)

	var topicRequest TopicRequest
	json.NewDecoder(r.Body).Decode(&topicRequest)
	defer r.Body.Close()

	topic := topicRequest.RequestToTopic(userID)

	// db.
	err := t.service.AddTopic(&topic)

	if err != nil {
		return nil, NewAppError(err)
	}

	// response.
	data := newTopicResponse(topic)
	return data, nil
}

func (t *TopicHandler) RemoveTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	// request.
	var body map[string]int
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()
	id, ok := body["id"]
	if !ok {
		e := errors.New("not id")
		return nil, NewAppError(e)
	}

	topicID := uint(id)
	id1, _ := strconv.Atoi(p.ByName("userID"))
	userID := uint(id1)

	// db.
	err := t.service.RemoveTopic(userID, topicID)
	if err != nil {
		return nil, NewAppError(err)
	}

	// response.
	return nil, nil
}

func (t *TopicHandler) UpdateTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	// request.
	id1, _ := strconv.Atoi(p.ByName("userID"))
	userID := uint(id1)

	var topicRequest TopicRequest
	json.NewDecoder(r.Body).Decode(&topicRequest)
	defer r.Body.Close()

	topic := topicRequest.RequestToTopic(userID)

	// db.
	updated, err := t.service.UpdateTopic(topic)
	if err != nil {
		return nil, NewAppError(err)
	}

	// response.
	data := newTopicResponse(*updated)
	return data, nil
}
