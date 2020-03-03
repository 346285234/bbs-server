package router

import (
	"encoding/json"
	"github.com/346285234/bbs-server/pkg/bbs"
	"github.com/346285234/bbs-server/pkg/user"
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

	// Get users.
	ids := make([]uint, len(favorites))
	i := 0
	for _, v := range favorites {
		ids[0] = v.UserID
		i++
	}
	users, _ := user.GetUsers(ids)

	userResponse := make([]UserResponse, len(users))
	for _, v := range users {
		r := UserResponse{v.ID, v.Name, v.Portrait}
		userResponse = append(userResponse, r)
	}

	var data = struct {
		Users []UserResponse `json:"users"`
	}{userResponse}

	return data, nil
}
