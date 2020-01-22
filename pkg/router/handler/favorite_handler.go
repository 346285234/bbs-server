package handler

import (
	"encoding/json"
	"net/http"

	"github.com/346285234/bbs-server/pkg/router"

	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/pkg/models"
	"github.com/346285234/bbs-server/pkg/services"
	"github.com/julienschmidt/httprouter"
)

type favoriteHanlder struct {
}

var FaH = favoriteHanlder{}

func (_ *favoriteHanlder) MarkFavorite(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *router.AppError) {

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
		return nil, router.NewAppError(err)
	}

	return nil, nil
}

func (_ *favoriteHanlder) CheckFavorite(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *router.AppError) {
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
