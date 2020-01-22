package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/346285234/bbs-server/pkg/models"
)

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

func TopicToResponse(topic models.Topic) router.TopicResponse {
	strings := TagsToStrings(topic.Tags)
	response := router.TopicResponse{topic.ID, topic.Title, strings,
		topic.CategoryID, topic.Category.Value,
		topic.UserID, topic.Intro, topic.Content,
		topic.UpdatedAt, topic.FavoritesCount,
		topic.LikeCount, topic.ViewCount}

	return response
}

func RequestToTopic(topicRequest router.TopicRequest, userID uint) models.Topic {
	tags := StringsToTags(topicRequest.Tags, userID)
	topic := models.Topic{Title: topicRequest.Title, Content: topicRequest.Content,
		CategoryID: topicRequest.CategoryID, Tags: tags,
		EditTime: topicRequest.EditTime, IsPaste: topicRequest.IsPaste,
		EditType: topicRequest.EditType, GroupID: topicRequest.GroupID,
		UserID: userID,
	}
	if topicRequest.ID != 0 {
		topic.ID = topicRequest.ID
	}

	return topic
}

func StringsToTags(strings []string, userID uint) []*models.Tag {
	tags := make([]*models.Tag, len(strings))
	for i, v := range strings {
		tags[i] = &models.Tag{UserID: userID, Value: v}
	}
	return tags
}

func TagsToStrings(tags []*models.Tag) []string {
	result := make([]string, len(tags))
	for i, v := range tags {
		result[i] = v.Value
	}

	return result
}

func CategoriesToResponse(categories []models.Category) []router.CategoryResponse {
	result := make([]router.CategoryResponse, len(categories))
	for i, v := range categories {
		result[i] = router.CategoryResponse{v.ID, v.Value}
	}
	return result
}

func CommentToResponse(comment models.Comment) router.CommentResponse {
	subComments := make([]router.CommentResponse, len(comment.Subs))
	for i, v := range comment.Subs {
		subComments[i] = router.CommentResponse{v.ID, v.AuthorID,
			v.Content, v.UpdatedAt, v.LikeCount,
			[]router.CommentResponse{}}
	}

	return router.CommentResponse{comment.ID, comment.AuthorID,
		comment.Content, comment.UpdatedAt, comment.LikeCount,
		subComments}
}

func CheckUser(r *http.Request) bool {
	return true
}

func Intro(from string) string {
	return from[:1]
}

func StrToInt(from string) int {
	result, error := strconv.Atoi(from)
	if error != nil {
		result = 0
	}

	return result
}
