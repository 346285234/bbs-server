package main

import (
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
	writer = httptest.NewRecorder()
}

func TestListTopic(t *testing.T) {
	request, _ := http.NewRequest("GET", "/topics", nil)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

}