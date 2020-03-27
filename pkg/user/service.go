package user

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type UserResponse struct {
	Success bool
	Code    int
	Message string
	Data    []User
}

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"username"`
	Portrait string `json:"avatar"`
}

func GetUsers(ids []uint) ([]User, error) {
	url := "http://localhost:8201/inner/users"
	contentType := "application/json"
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(ids)

	resp, err := http.Post(url, contentType, buf)
	if err != nil {
		log.Printf("fetch: %v\n", err)
		return nil, err
	}

	var r UserResponse
	json.NewDecoder(resp.Body).Decode(&r)

	if r.Success == true {
		return r.Data, nil
	} else {
		// TODO: return error base on message from user server.
		return nil, nil
	}
}

func GetUser(id uint) (*User, error) {
	users, err := GetUsers([]uint{id})
	if err != nil {
		return nil, err
	}

	return &users[0], nil
}
