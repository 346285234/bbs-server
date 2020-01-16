package handler

import (
	"encoding/json"
	"errors"
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type topicHandler struct {
}

var Th = topicHandler{}

func (_ *topicHandler) ListTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {
	// TODO: request.

	// db.
	topics, err := services.Ts.Topics()
	if err != nil {
		return nil, models.NewAppError(err)
	}

	// response.
	topicsResponse := make([]models.TopicResponse, len(topics))
	for i, v := range topics {
		topicResponse := common.TopicToResponse(v)
		topicsResponse[i] = topicResponse
	}

	data := struct {
		Total  int                    `json:"total"`
		Topics []models.TopicResponse `json:"topics"`
	}{len(topicsResponse), topicsResponse}
	return data, nil
}

func (_ *topicHandler) GetTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {

	// request.
	id := uint(common.StrToInt(p.ByName("id")))

	// db.
	topic, err := services.Ts.GetTopic(id)
	if err != nil {
		return nil, models.NewAppError(err)
	}

	// response.
	topicResponse := common.TopicToResponse(*topic)
	return topicResponse, nil
}

func (_ *topicHandler) AddTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {
	// request.
	userID := uint(common.StrToInt(r.Header.Get("userID")))

	var topicRequest models.TopicRequest
	json.NewDecoder(r.Body).Decode(&topicRequest)
	defer r.Body.Close()

	topic := common.RequestToTopic(topicRequest)
	topic.UserID = userID

	// db.
	err := services.Ts.AddTopic(&topic)

	if err != nil {
		return nil, models.NewAppError(err)
	}

	// response.
	data := common.TopicToResponse(topic)
	return data, nil
}

func (_ *topicHandler) RemoveTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {
	// request.
	var body map[string]int
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()
	id, ok := body["id"]
	if !ok {
		e := errors.New("not id")
		return nil, models.NewAppError(e)
	}

	topicID := uint(id)
	userID := uint(common.StrToInt(r.Header.Get("userID")))

	// db.
	err := services.Ts.RemoveTopic(userID, topicID)
	if err != nil {
		return nil, models.NewAppError(err)
	}

	// response.
	return nil, nil
}

func (_ *topicHandler) UpdateTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {
	// request.
	userID := uint(common.StrToInt(r.Header.Get("userID")))

	var topicRequest models.TopicRequest
	json.NewDecoder(r.Body).Decode(&topicRequest)
	defer r.Body.Close()

	topic := common.RequestToTopic(topicRequest)
	topic.UserID = userID

	// db.
	err := services.Ts.UpdateTopic(&topic)
	if err != nil {
		return nil, models.NewAppError(err)
	}

	// response.
	data := common.TopicToResponse(topic)
	return data, nil
}
