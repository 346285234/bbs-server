package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/346285234/bbs-server/pkg/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/346285234/bbs-server/pkg/router"
	"github.com/julienschmidt/httprouter"
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
	db := gorm.Open("mysql", "root:346285234@tcp(localhost)/bbs?parseTime=true")

	categoryService := gorm.NewCategoryService(db)
	commentService := gorm.NewCommentService(db)
	favoriteService := gorm.NewFavoriteService(db)
	likeService := gorm.NewLikeService(db)
	tagService := gorm.NewTagService(db)
	topicService := gorm.NewTopicService(db)

	categoryHandler := router.NewCategoryHandler(&categoryService)
	commentHandler := router.NewCommentHandler(&commentService)
	favoriteHandler := router.NewFavoriteHandler(&favoriteService)
	likeHandler := router.NewLikeHandler(&likeService)
	tagHandler := router.NewTagHandler(&tagService)
	topicHandler := router.NewTopicHandler(&topicService)

	handlers := []interface{}{categoryHandler, commentHandler, favoriteHandler, likeHandler,
		tagHandler, topicHandler}
	r = router.NewRouter(handlers)
	writer = httptest.NewRecorder()
	flag.Parse()
}

// MARK: Topic.

func TestListTopic(t *testing.T) {
	request, _ := http.NewRequest("GET", "/topics?tag=h&page=1&page_size=1000", nil)
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	fmt.Println(writer.Body)

}

func TestGetTopic(t *testing.T) {
	request, _ := http.NewRequest("GET", "/topic/2", nil)
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
		Title:      "third topic",
		Content:    "hello world!",
		CategoryID: 2,
		Tags:       []string{"a", "h"},
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
	body := map[string]int{"id": 3}
	json.NewEncoder(buf).Encode(&body)
	request, _ := http.NewRequest("POST", "/topic/remove", buf)

	request.Header.Add("userID", string(userID))
	r.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	fmt.Println(writer.Body)
}

func TestUpdateTopic(t *testing.T) {
	buf := new(bytes.Buffer)
	type TopicRequest struct {
		ID         uint
		Title      string
		Content    string
		CategoryID uint `json:"category_id"`
		Tags       []string
		EditTime   time.Duration `json:"edit_time"`
		IsPaste    bool          `json:"is_paste"`
		EditType   uint          `json:"edit_type"`
		GroupID    uint          `json:"group_id"`
	}

	body := &TopicRequest{
		ID:         2,
		Title:      "update first",
		Content:    "hello",
		CategoryID: 3,
		Tags:       []string{"a", "g"},
		EditTime:   time.Minute,
		IsPaste:    true,
		EditType:   1,
		GroupID:    10,
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
	var body = &RequestBody{false}
	json.NewEncoder(buf).Encode(&body)
	request, _ := http.NewRequest("POST", "/favorite/topic/2/mark", buf)
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
