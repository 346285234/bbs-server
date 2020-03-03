package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	var r UserResponse
	json.NewDecoder(resp.Body).Decode(&r)

	if r.Success == true {
		return r.Data, nil
	} else {
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
