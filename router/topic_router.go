package router

import (
	"encoding/json"
	"errors"
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/services"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type TopicRouter struct {
	logger log.Logger
}

var tr = TopicRouter{}

func (_ *TopicRouter) SetupRouter(mux http.ServeMux) {
}

func (_ *TopicRouter) listTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) *appError {
	// TODO: Analyse query.

	// Get data.
	topics, err := services.Ts.Topics()

	if err != nil {
		return NewAppError(err)
	}

	// Set response.
	var response Response
	response = Response{Success: true, Code: 200, Message: "OK"}
	data := struct {
		Total  int            `json:"total"`
		Topics []models.Topic `json:"topics"`
	}{len(topics), topics}
	response.Data = data

	bytes, err := json.Marshal(response)

	if err != nil {
		return NewAppError(err)
	}

	w.Write(bytes)

	return nil
}

func (_ *TopicRouter) getTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) *appError {

	// Get id.
	id := common.StrToInt(p.ByName("id"))

	// Get data.
	topic, err := services.Ts.GetTopic(uint(id))

	if err != nil {
		return NewAppError(err)
	}
	// Set response.
	var response Response
	response = Response{Success: true, Code: 200, Message: "OK"}
	response.Data = topic

	bytes, err := json.Marshal(response)
	if err != nil {
		return NewAppError(err)
	}

	w.Write(bytes)

	return nil
}

func (_ *TopicRouter) addTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) *appError {
	// Analyse response.

	var topic models.Topic
	json.NewDecoder(r.Body).Decode(&topic)
	defer r.Body.Close()
	userID := r.Header.Get("userID")
	topic.UserID = uint(common.StrToInt(userID))

	//Add data.
	err := services.Ts.AddTopic(topic)

	if err != nil {
		return NewAppError(err)
	}

	// Set response.
	var response Response
	response = Response{Success: true, Code: 200, Message: "OK"}

	bytes, err := json.Marshal(response)
	if err != nil {
		return NewAppError(err)
	}

	w.Write(bytes)
	return nil
}

func (_ *TopicRouter) removeTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) *appError {
	// Get id.
	var body map[string]int
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()
	id, ok := body["id"]
	if !ok {
		e := errors.New("not id")
		return NewAppError(e)
	}

	topicID := uint(id)
	userID := uint(common.StrToInt(r.Header.Get("userID")))

	err := services.Ts.RemoveTopic(userID, topicID)

	if err != nil {
		return NewAppError(err)
	}

	var response Response
	response = Response{Success: true, Code: 200, Message: "OK"}

	bytes, err := json.Marshal(response)
	if err != nil {
		return NewAppError(err)
	}

	w.Write(bytes)
	return nil
}

func (_ *TopicRouter) updateTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) *appError {
	var topic models.Topic
	json.NewDecoder(r.Body).Decode(&topic)
	defer r.Body.Close()
	userID := r.Header.Get("userID")
	topic.UserID = uint(common.StrToInt(userID))

	err := services.Ts.AddTopic(topic)
	if err != nil {
		return NewAppError(err)
	}

	var response Response
	response = Response{Success: true, Code: 200, Message: "OK"}

	bytes, err := json.Marshal(response)
	if err != nil {
		return NewAppError(err)
	}
	w.Write(bytes)

	return nil
}


func (_ *TopicRouter) listCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params) *appError {
	// Get data.
	categories, err := services.Cs.Categories()

	if err != nil {
		return NewAppError(err)
	}

	// Set response.
	var response Response
	response = Response{Success: true, Code: 200, Message: "OK"}
	data := struct {
		Total  int
		Categories []models.Category
	}{len(categories), categories}
	response.Data = data

	bytes, err := json.Marshal(response)

	if err != nil {
		return NewAppError(err)
	}

	w.Write(bytes)

	return nil
}

func (_ *TopicRouter) listTag(w http.ResponseWriter, r *http.Request, p httprouter.Params) *appError {
	// Get data.
	tags, err := services.TagS.Tags()

	if err != nil {
		return NewAppError(err)
	}

	// Set response.
	var response Response
	response = Response{Success: true, Code: 200, Message: "OK"}
	data := struct {
		Total  int
		Tags []models.Tag
	}{len(tags), tags}
	response.Data = data

	bytes, err := json.Marshal(response)

	if err != nil {
		return NewAppError(err)
	}

	w.Write(bytes)

	return nil
}

func (_ *TopicRouter) markFavorite(w http.ResponseWriter, r *http.Request, p httprouter.Params) *appError {

	userIDStr := r.Header.Get("userID")
	userID := uint(common.StrToInt(userIDStr))

	type RequestBody struct {
		topicID uint
		isMark bool `json:"type"`
	}
	var body RequestBody
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	var favorite models.TopicFavorite
	favorite.TopicID = body.topicID
	favorite.UserID = userID

	// Get data.
	err := services.Fs.Mark(favorite, body.isMark)

	if err != nil {
		return NewAppError(err)
	}

	// Set response.
	var response Response
	response = Response{Success: true, Code: 200, Message: "OK"}

	bytes, err := json.Marshal(response)

	if err != nil {
		return NewAppError(err)
	}

	w.Write(bytes)

	return nil
}

func (_ *TopicRouter) markLike(w http.ResponseWriter, r *http.Request, p httprouter.Params) *appError {

	userIDStr := r.Header.Get("userID")
	userID := uint(common.StrToInt(userIDStr))

	type RequestBody struct {
		topicID uint
		isMark bool `json:"type"`
	}
	var body RequestBody
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	var like models.TopicLike
	like.TopicID = body.topicID
	like.UserID = userID

	// Get data.
	err := services.Ls.Mark(like, body.isMark)

	if err != nil {
		return NewAppError(err)
	}

	// Set response.
	var response Response
	response = Response{Success: true, Code: 200, Message: "OK"}

	bytes, err := json.Marshal(response)

	if err != nil {
		return NewAppError(err)
	}

	w.Write(bytes)

	return nil
}