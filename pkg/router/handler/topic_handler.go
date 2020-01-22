package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/346285234/bbs-server/pkg/models"

	"github.com/346285234/bbs-server/pkg/router"

	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/pkg/services"
	"github.com/julienschmidt/httprouter"
)

type topicHandler struct {
	service models.TopicService
}

var Th = topicHandler{}

func (_ *topicHandler) ListTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *router.AppError) {
	// TODO: request.

	vars := r.URL.Query()

	query := make(map[string]interface{})
	groupID := uint(common.StrToInt(vars.Get("group_id")))
	if groupID != 0 {
		query["group_id"] = groupID
	}
	userID := uint(common.StrToInt(vars.Get("user_id")))
	if userID != 0 {
		query["user_id"] = userID
	}
	categoryID := uint(common.StrToInt(vars.Get("category_id")))
	if categoryID != 0 {
		query["category_id"] = categoryID
	}
	page := uint(common.StrToInt(vars.Get("page")))
	pageSize := uint(common.StrToInt(vars.Get("page_size")))
	if page != 0 && pageSize != 0 {
		query["page"] = page
		query["page_size"] = pageSize
	}

	tag := vars.Get("tag")
	if tag != "" {
		query["tag"] = tag
	}

	// db.
	topics, err := services.Ts.Topics(query)
	if err != nil {
		return nil, router.NewAppError(err)
	}

	// response.
	topicsResponse := make([]router.TopicResponse, len(topics))
	for i, v := range topics {
		topicResponse := common.TopicToResponse(v)
		topicsResponse[i] = topicResponse
	}

	data := struct {
		Total  int                    `json:"total"`
		Topics []router.TopicResponse `json:"topics"`
	}{len(topicsResponse), topicsResponse}
	return data, nil
}

func (_ *topicHandler) GetTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *router.AppError) {

	// request.
	id := uint(common.StrToInt(p.ByName("id")))

	// db.
	topic, err := services.Ts.GetTopic(id)
	if err != nil {
		return nil, router.NewAppError(err)
	}

	// response.
	topicResponse := common.TopicToResponse(*topic)
	return topicResponse, nil
}

func (_ *topicHandler) AddTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *router.AppError) {
	// request.
	userID := uint(common.StrToInt(r.Header.Get("userID")))

	var topicRequest router.TopicRequest
	json.NewDecoder(r.Body).Decode(&topicRequest)
	defer r.Body.Close()

	topic := common.RequestToTopic(topicRequest, userID)

	// db.
	err := services.Ts.AddTopic(&topic)

	if err != nil {
		return nil, router.NewAppError(err)
	}

	// response.
	data := common.TopicToResponse(topic)
	return data, nil
}

func (_ *topicHandler) RemoveTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *router.AppError) {
	// request.
	var body map[string]int
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()
	id, ok := body["id"]
	if !ok {
		e := errors.New("not id")
		return nil, router.NewAppError(e)
	}

	topicID := uint(id)
	userID := uint(common.StrToInt(r.Header.Get("userID")))

	// db.
	err := services.Ts.RemoveTopic(userID, topicID)
	if err != nil {
		return nil, router.NewAppError(err)
	}

	// response.
	return nil, nil
}

func (_ *topicHandler) UpdateTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *router.AppError) {
	// request.
	userID := uint(common.StrToInt(r.Header.Get("userID")))

	var topicRequest router.TopicRequest
	json.NewDecoder(r.Body).Decode(&topicRequest)
	defer r.Body.Close()

	topic := common.RequestToTopic(topicRequest, userID)

	// db.
	updated, err := services.Ts.UpdateTopic(topic)
	if err != nil {
		return nil, router.NewAppError(err)
	}

	// response.
	data := common.TopicToResponse(*updated)
	return data, nil
}
