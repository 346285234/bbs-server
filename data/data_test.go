package data

import "testing"

func TestCreateTable(t *testing.T) {
	models := []interface{}{&Topic{}, &TopicFavorite{}, &TopicLike{},
		&Category{}, &Comment{}, &CommentLike{}}

	for _, v := range models {
		Db.CreateTable(v)
	}
}