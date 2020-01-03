package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/346285234/bbs-server/router"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var r *httprouter.Router
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M)  {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	r = router.NewRouter()
	writer = httptest.NewRecorder()
}

func TestListTopic(t *testing.T) {
	request, _ := http.NewRequest("GET", "/topics", nil)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)

}

func TestGetTopic(t *testing.T) {
	request, _ := http.NewRequest("GET", "/topic/1", nil)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)
}

func TestAddTopic(t *testing.T) {
	buf := new(bytes.Buffer)
	type RequestBody struct {
		Title string
		Content string
		CategoryID uint `json:"category_id"`
		Tags []string `json:"tags"` // FIXME: get tags.
		EditTime time.Duration `json:"edit_time"`
		IsPaste bool `json:"is_paste"`
		EditType int `json:"edit_type"`
		GroupID int `json:"group_id"`
	}
	body := &RequestBody{
		Title:   "second",
		Content: "hello world",
		CategoryID: 1,
		Tags: []string{"a", "b"},
		EditTime: time.Hour,
		IsPaste: false,
		EditType: 1,
		GroupID: 1,
	}
	json.NewEncoder(buf).Encode(&body)

	request, _ := http.NewRequest("POST", "/topic/add", buf)

	request.Header.Add("userID", "1")
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	fmt.Println(writer.Body)
}

func TestRemoveTopic(t *testing.T) {
	buf := new(bytes.Buffer)
	body := map[string]int{"id":2}
	json.NewEncoder(buf).Encode(&body)
	request, _ := http.NewRequest("POST", "/topic/remove", buf)

	request.Header.Add("userID", "1")
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	fmt.Println(writer.Body)
}

func TestUpdateTopic(t *testing.T) {
	buf := new(bytes.Buffer)
	type Body struct {
		ID int `json:"id"`
		Title string `json:"title""`
		Content string `json:"content"`
	}
	body := &Body{
		ID: 1,
		Title:   "first",
		Content: "update content",
	}
	json.NewEncoder(buf).Encode(&body)

	request, _ := http.NewRequest("POST", "/topic/update", buf)
	request.Header.Add("userID", "1")
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	fmt.Println(writer.Body)
}