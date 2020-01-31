package router

import (
	"encoding/json"
	"github.com/346285234/bbs-server/pkg/bbs"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type FavoriteHanlder struct {
	service bbs.FavoriteService
}

func NewFavoriteHandler(s bbs.FavoriteService) FavoriteHanlder {
	return FavoriteHanlder{s}
}

func (f *FavoriteHanlder) MarkFavorite(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	id1, _ := strconv.Atoi(p.ByName("topic_id"))
	id2, _ := strconv.Atoi(r.Header.Get("userID"))
	topicID := uint(id1)
	userID := uint(id2)
	var favorite bbs.Favorite
	favorite.TopicID = topicID
	favorite.UserID = userID

	type RequestBody struct {
		IsMark bool `json:"is_mark"`
	}
	var body RequestBody
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	// Get data.
	err := f.service.Mark(favorite, body.IsMark)

	if err != nil {
		return nil, NewAppError(err)
	}

	return nil, nil
}

func (f *FavoriteHanlder) CheckFavorite(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	id1, _ := strconv.Atoi(p.ByName("topic_id"))
	id2, _ := strconv.Atoi(r.Header.Get("userID"))
	topicID := uint(id1)
	userID := uint(id2)

	var favorite bbs.Favorite
	favorite.TopicID = topicID
	favorite.UserID = userID

	isMark, _ := f.service.Check(favorite)
	var data = struct {
		IsMark bool `json:"is_mark"`
	}{isMark}

	return data, nil
}

func (f *FavoriteHanlder) FavoriteUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	id1, _ := strconv.Atoi(p.ByName("topic_id"))
	topicID := uint(id1)
	favorites, _ := f.service.List(topicID)

	// TODO: Get users.
	var users = make([]User, len(favorites))
	for i, v := range favorites {
		users[i] = User{ID: v.UserID}
	}

	var data = struct {
		Users []User `json:"users"`
	}{users}

	return data, nil
}
