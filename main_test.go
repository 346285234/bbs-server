package main

import (
	"bytes"
	"encoding/json"
	"github.com/346285234/bbs-server/router"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var r *httprouter.Router
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M)  {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	r = httprouter.New()
	r.GET("/topics", router.Tr.ListTopics)
	r.GET("/topic/:id", router.Tr.GetTopic)
	r.POST("/topic/add", check(router.Tr.CreateTopic))
	writer = httptest.NewRecorder()
}

func TestListTopic(t *testing.T) {
	request, _ := http.NewRequest("GET", "/topics", nil)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

}

func TestGetTopic(t *testing.T) {
	request, _ := http.NewRequest("GET", "/topic/1", nil)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

}

func TestCreateTopic(t *testing.T) {
	type CreateTopic struct {
		Title string `json:"title""`
		Content string `json:"content"`
	}
	ct := &CreateTopic{
		Title:   "first",
		Content: "hello world",
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&ct)
	request, _ := http.NewRequest("POST", "/topic/add", buf)
	request.Header.Add("userID", "1")
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

}