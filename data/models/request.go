package models

import "time"

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
