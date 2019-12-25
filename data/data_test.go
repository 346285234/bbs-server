package data

import "testing"

func TestCreateTable(t *testing.T) {
	//models := []interface{}{&Topic{}, &TopicFavorite{}, &TopicLike{},
	//	&Category{}, &Comment{}, &CommentLike{}}

	models := []interface{}{&Topic{}}
	for _, v := range models {
		db.CreateTable(v)
	}
}