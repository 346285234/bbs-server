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
	// TODO: Analyse query.

	// Get data.
	topics, err := services.Ts.Topics()

	if err != nil {
		return nil, models.NewAppError(err)
	}

	data := struct {
		Total  int            `json:"total"`
		Topics []models.Topic `json:"topics"`
	}{len(topics), topics}
	return data, nil
}

func (_ *topicHandler) GetTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {

	// Get id.
	id := uint(common.StrToInt(p.ByName("id")))

	// Get data.
	topic, err := services.Ts.GetTopic(id)

	if err != nil {
		return nil, models.NewAppError(err)
	}

	return topic, nil
}

func (_ *topicHandler) AddTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {
	// Analyse response.

	var topic models.Topic
	json.NewDecoder(r.Body).Decode(&topic)
	defer r.Body.Close()
	userID := uint(common.StrToInt(r.Header.Get("userID")))
	topic.UserID = userID

	//Add data.
	err := services.Ts.AddTopic(topic)

	if err != nil {
		return nil, models.NewAppError(err)
	}

	return nil, nil
}

func (_ *topicHandler) RemoveTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {
	// Get id.
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

	err := services.Ts.RemoveTopic(userID, topicID)

	if err != nil {
		return nil, models.NewAppError(err)
	}

	return nil, nil
}

func (_ *topicHandler) UpdateTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{},*models.AppError) {
	var topic models.Topic
	json.NewDecoder(r.Body).Decode(&topic)
	defer r.Body.Close()
	userID := r.Header.Get("userID")
	topic.UserID = uint(common.StrToInt(userID))

	err := services.Ts.AddTopic(topic)
	if err != nil {
		return nil, models.NewAppError(err)
	}

	return nil, nil
}

