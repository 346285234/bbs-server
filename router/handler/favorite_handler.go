package handler

import (
	"encoding/json"
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data/models"
	"github.com/346285234/bbs-server/data/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type favoriteHanlder struct {
}

var FaH = favoriteHanlder{}

func (_ *favoriteHanlder) MarkFavorite(w http.ResponseWriter, r *http.Request, p httprouter.Params) *models.AppError {

	topicID := uint(common.StrToInt(p.ByName("topic_id")))
	userID := uint(common.StrToInt(r.Header.Get("userID")))
	var favorite models.TopicFavorite
	favorite.TopicID = topicID
	favorite.UserID = userID

	type RequestBody struct {
		Unmark bool `json:"unmark"`
	}
	var body RequestBody
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	// Get data.
	err := services.Fs.Mark(favorite, !body.Unmark)

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

func (_ *favoriteHanlder) CheckFavorite(w http.ResponseWriter, r *http.Request, p httprouter.Params) *models.AppError {
	topicID := uint(common.StrToInt(p.ByName("topic_id")))
	userID := uint(common.StrToInt(r.Header.Get("userID")))

	var favorite models.TopicFavorite
	favorite.TopicID = topicID
	favorite.UserID = userID

	err := services.Fs.Check(favorite)
	var isMark bool
	if err == nil {
		isMark = true
	} else {
		isMark = false
	}
	var data = struct{
		IsMark bool `json:"is_mark"`
	}{isMark}

	// Set response.
	var response models.Response
	response = models.Response{true, 200, "OK", data}
	bytes, err := json.Marshal(response)
	if err != nil {
		return models.NewAppError(err)
	}

	w.Write(bytes)

	return nil
}