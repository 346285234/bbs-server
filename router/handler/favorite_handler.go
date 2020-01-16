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

func (_ *favoriteHanlder) MarkFavorite(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {

	topicID := uint(common.StrToInt(p.ByName("topic_id")))
	userID := uint(common.StrToInt(r.Header.Get("userID")))
	var favorite models.Favorite
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
		return nil, models.NewAppError(err)
	}

	return nil, nil
}

func (_ *favoriteHanlder) CheckFavorite(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *models.AppError) {
	topicID := uint(common.StrToInt(p.ByName("topic_id")))
	userID := uint(common.StrToInt(r.Header.Get("userID")))

	var favorite models.Favorite
	favorite.TopicID = topicID
	favorite.UserID = userID

	isMark, _ := services.Fs.Check(favorite)
	var data = struct {
		IsMark bool `json:"is_mark"`
	}{isMark}

	return data, nil
}
