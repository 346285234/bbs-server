package common

import (
	"github.com/346285234/bbs-server/data/models"
	"net/http"
	"strconv"
)

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

func TopicToResponse(topic models.Topic) models.TopicResponse {
	strings := TagsToStrings(topic.Tags)
	response := models.TopicResponse{topic.ID, topic.Title, strings,
		topic.CategoryID, topic.Category.Value,
		topic.UserID, topic.Intro, topic.Content,
		topic.UpdatedAt, topic.FavoritesCount,
		topic.LikeCount, topic.ViewCount}

	return response
}

func RequestToTopic(topicRequest models.TopicRequest) models.Topic {
	tags := StringsToTags(topicRequest.Tags)
	topic := models.Topic{Title: topicRequest.Title, Content: topicRequest.Content,
		CategoryID: topicRequest.CategoryID, Tags: tags,
		EditTime: topicRequest.EditTime, IsPaste: topicRequest.IsPaste,
		EditType: topicRequest.EditType, GroupID: topicRequest.GroupID,
	}
	return topic
}

func StringsToTags(strings []string) []*models.Tag {
	tags := make([]*models.Tag, len(strings))
	for i, v := range strings {
		tags[i] = &models.Tag{Value: v}
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

func CategoriesToResponse(categories []models.Category) []models.CategoryResponse {
	result := make([]models.CategoryResponse, len(categories))
	for i, v := range categories {
		result[i] = models.CategoryResponse{v.ID, v.Value}
	}
	return result
}

func CommentToResponse(comment models.Comment) models.CommentResponse {
	subComments := make([]models.CommentResponse, len(comment.Subs))
	for i, v := range comment.Subs {
		subComments[i] = models.CommentResponse{v.ID, v.AuthorID,
			v.Content, v.UpdatedAt, v.LikeCount,
			[]models.CommentResponse{}}
	}

	return models.CommentResponse{comment.ID, comment.AuthorID,
		comment.Content, comment.UpdatedAt, comment.LikeCount,
		subComments}
}
