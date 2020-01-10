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

func (_ *topicHandler) ListTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) *models.AppError {
	// TODO: Analyse query.

	// Get data.
	topics, err := services.Ts.Topics()

	if err != nil {
		return models.NewAppError(err)
	}

	// Set response.
	var response models.Response
	response = models.Response{Success: true, Code: 200, Message: "OK"}
	data := struct {
		Total  int            `json:"total"`
		Topics []models.Topic `json:"topics"`
	}{len(topics), topics}
	response.Data = data

	bytes, err := json.Marshal(response)

	if err != nil {
		return models.NewAppError(err)
	}

	w.Write(bytes)

	return nil
}

func (_ *topicHandler) GetTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) *models.AppError {

	// Get id.
	id := common.StrToInt(p.ByName("id"))

	// Get data.
	topic, err := services.Ts.GetTopic(uint(id))

	if err != nil {
		return models.NewAppError(err)
	}
	// Set response.
	var response models.Response
	response = models.Response{Success: true, Code: 200, Message: "OK"}
	response.Data = topic

	bytes, err := json.Marshal(response)
	if err != nil {
		return models.NewAppError(err)
	}

	w.Write(bytes)

	return nil
}

func (_ *topicHandler) AddTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) *models.AppError {
	// Analyse response.

	var topic models.Topic
	json.NewDecoder(r.Body).Decode(&topic)
	defer r.Body.Close()
	userID := r.Header.Get("userID")
	topic.UserID = uint(common.StrToInt(userID))

	//Add data.
	err := services.Ts.AddTopic(topic)

	if err != nil {
		return models.NewAppError(err)
	}

	// Set response.
	var response models.Response
	response = models.Response{Success: true, Code: 200, Message: "OK"}

	bytes, err := json.Marshal(response)
	if err != nil {
		return models.NewAppError(err)
	}

	w.Write(bytes)
	return nil
}

func (_ *topicHandler) RemoveTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) *models.AppError {
	// Get id.
	var body map[string]int
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()
	id, ok := body["id"]
	if !ok {
		e := errors.New("not id")
		return models.NewAppError(e)
	}

	topicID := uint(id)
	userID := uint(common.StrToInt(r.Header.Get("userID")))

	err := services.Ts.RemoveTopic(userID, topicID)

	if err != nil {
		return models.NewAppError(err)
	}

	var response models.Response
	response = models.Response{Success: true, Code: 200, Message: "OK"}

	bytes, err := json.Marshal(response)
	if err != nil {
		return models.NewAppError(err)
	}

	w.Write(bytes)
	return nil
}

func (_ *topicHandler) UpdateTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) *models.AppError {
	var topic models.Topic
	json.NewDecoder(r.Body).Decode(&topic)
	defer r.Body.Close()
	userID := r.Header.Get("userID")
	topic.UserID = uint(common.StrToInt(userID))

	err := services.Ts.AddTopic(topic)
	if err != nil {
		return models.NewAppError(err)
	}

	var response models.Response
	response = models.Response{Success: true, Code: 200, Message: "OK"}

	bytes, err := json.Marshal(response)
	if err != nil {
		return models.NewAppError(err)
	}
	w.Write(bytes)

	return nil
}

