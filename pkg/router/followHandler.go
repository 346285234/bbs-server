package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/346285234/bbs-server/pkg/bbs"
	"github.com/346285234/bbs-server/pkg/user"
	"github.com/julienschmidt/httprouter"
)

type FollowHanlder struct {
	service bbs.FollowService
}

func NewFollowHandler(s bbs.FollowService) FollowHanlder {
	return FollowHanlder{s}
}

func (f *FollowHanlder) Mark(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	id1, _ := strconv.Atoi(p.ByName("user_id"))
	id2, _ := strconv.Atoi(r.Header.Get("userID"))
	objectID := uint(id1)
	subjectID := uint(id2)
	var follow bbs.Follow
	follow.SubjectID = subjectID
	follow.ObjectID = objectID

	type RequestBody struct {
		IsMark bool `json:"is_mark"`
	}
	var body RequestBody
	json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	// Get data.
	err := f.service.Mark(follow, body.IsMark)

	if err != nil {
		return nil, NewAppError(err)
	}

	return nil, nil
}

func (f *FollowHanlder) Check(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	id1, _ := strconv.Atoi(p.ByName("user_id"))
	id2, _ := strconv.Atoi(r.Header.Get("userID"))
	objectID := uint(id1)
	subjectID := uint(id2)

	var follow bbs.Follow
	follow.SubjectID = subjectID
	follow.ObjectID = objectID

	isMark, _ := f.service.Check(follow)
	var data = struct {
		IsMark bool `json:"is_mark"`
	}{isMark}

	return data, nil
}

func (f *FollowHanlder) FollowUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	id1, _ := strconv.Atoi(p.ByName("user_id"))

	follows, _ := f.service.List(uint(id1))

	// Get users.
	ids := make([]uint, len(follows))
	i := 0
	for _, v := range follows {
		ids[i] = v.ObjectID
		i++
	}
	users, err := user.GetUsers(ids)
	if err != nil {
		return nil, NewAppError(err)
	}

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
