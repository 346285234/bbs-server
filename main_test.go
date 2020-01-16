package main

import (
	"bytes"
	"encoding/json"
	"flag"
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
var userID = "1"

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	r = router.NewRouter()
	writer = httptest.NewRecorder()
	flag.Parse()
}

// MARK: Topic.

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
		Title      string        `json:"title"`
		Content    string        `json:"content"`
		CategoryID uint          `json:"category_id"`
		Tags       []string      `json:"tags"`
		EditTime   time.Duration `json:"edit_time"`
		IsPaste    bool          `json:"is_paste"`
		EditType   uint          `json:"edit_type"`
		GroupID    uint          `json:"group_id"`
	}
	body := RequestBody{
		Title:      "fifth topic",
		Content:    "hello world!",
		CategoryID: 1,
		Tags:       []string{"a", "b"},
		EditTime:   time.Hour,
		IsPaste:    true,
		EditType:   1,
		GroupID:    1,
	}
	json.NewEncoder(buf).Encode(body)
	request, _ := http.NewRequest("POST", "/topic/add", buf)

	request.Header.Add("userID", userID)

	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	fmt.Println(writer.Body)
}

func TestRemoveTopic(t *testing.T) {
	buf := new(bytes.Buffer)
	body := map[string]int{"id": 5}
	json.NewEncoder(buf).Encode(&body)
	request, _ := http.NewRequest("POST", "/topic/remove", buf)

	request.Header.Add("userID", string(userID))
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	fmt.Println(writer.Body)
}

// TODO: update topic api.
func TestUpdateTopic(t *testing.T) {
	buf := new(bytes.Buffer)
	type Body struct {
		ID      int    `json:"id"`
		Title   string `json:"title""`
		Content string `json:"content"`
	}
	body := &Body{
		ID:      1,
		Title:   "first",
		Content: "update content",
	}
	json.NewEncoder(buf).Encode(&body)

	request, _ := http.NewRequest("POST", "/topic/update", buf)
	request.Header.Add("userID", string(userID))
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	fmt.Println(writer.Body)
}

// MARK: Category & Tag.

func TestListCategory(t *testing.T) {
	request, _ := http.NewRequest("GET", "/categories", nil)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)

}

func TestListTag(t *testing.T) {
	request, _ := http.NewRequest("GET", "/tags", nil)
	request.Header.Set("userID", string(userID))
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)

}

// MARK: Like & favorite.

func TestMarkFavorite(t *testing.T) {
	buf := new(bytes.Buffer)
	type RequestBody struct {
		Unmark bool `json:"unmark"`
	}
	var body = &RequestBody{true}
	json.NewEncoder(buf).Encode(&body)
	request, _ := http.NewRequest("POST", "/favorite/topic/1/mark", buf)
	request.Header.Add("userID", userID)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)
}

func TestCheckFavorite(t *testing.T) {
	request, _ := http.NewRequest("GET", "/favorite/topic/1", nil)
	request.Header.Add("userID", userID)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)

}

func TestMarkLikeTopic(t *testing.T) {
	buf := new(bytes.Buffer)
	type RequestBody struct {
		Unmark bool `json:"unmark"`
	}
	var body = &RequestBody{true}
	json.NewEncoder(buf).Encode(&body)
	request, _ := http.NewRequest("POST", "/like/topic/1/mark", buf)
	request.Header.Add("userID", userID)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)
}

func TestCheckLikeTopic(t *testing.T) {
	request, _ := http.NewRequest("GET", "/like/topic/1", nil)
	request.Header.Add("userID", userID)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)

}

func TestMarkLikeComment(t *testing.T) {
	buf := new(bytes.Buffer)
	type RequestBody struct {
		Unmark bool `json:"unmark"`
	}
	var body = &RequestBody{true}
	json.NewEncoder(buf).Encode(&body)
	request, _ := http.NewRequest("POST", "/like/comment/2/mark", buf)
	request.Header.Add("userID", userID)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)
}

func TestCheckLikeComment(t *testing.T) {
	request, _ := http.NewRequest("GET", "/like/comment/2", nil)
	request.Header.Add("userID", userID)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)

}

// MARK: Comment

func TestListComment(t *testing.T) {
	request, _ := http.NewRequest("GET", "/comments/1", nil)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)

}

func TestReplyComment(t *testing.T) {
	comment := struct {
		ParentID int `json:"parent_id"`
		Content  string
	}{0, "fourth comment"}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&comment)

	request, _ := http.NewRequest("POST", "/comment/1/reply", buf)
	request.Header.Add("userID", userID)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)

}

// TODO: undo.
func TestRevokeComment(t *testing.T) {
	request, _ := http.NewRequest("POST", "/comment/1/revoke", nil)
	request.Header.Add("userID", userID)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)

}
