package models

import "time"

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type TopicResponse struct {
	ID           uint     `json:"id"`
	Title        string   `json:"title"`
	Tags         []string `json:"tags"`
	CategoryID   uint     `json:"category_id"`
	CategoryName string   `json:"category_name"`
	AuthorID     uint     `json:"author_id"`
	//AuthorName string `json:"author_name"`
	//AuthorPortrait string `json:"author_portrait"`
	Description   string    `json:"description"`
	Content       string    `json:"content"`
	ModifyTime    time.Time `json:"modify_time"`
	FavoriteCount uint      `json:"favorite_count"`
	LikeCount     uint      `json:"like_count"`
	ViewCount     uint      `json:"view_count"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CommentResponse struct {
	ID       uint `json:"id"`
	AuthorID uint `json:"author_id"`
	//AuthorName string `json:"author_name"`
	//AuthorPortrait string `json:"author_portrait"`
	Content    string            `json:"content"`
	ModifyTime time.Time         `json:"modify_time"`
	LikeCount  uint              `json:"like_count"`
	Subs       []CommentResponse `json:"sub_comments"`
}
