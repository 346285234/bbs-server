package router

import (
	"github.com/346285234/bbs-server/pkg/bbs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type tags []string

func newTags(tags []*bbs.Tag) []string {
	result := make([]string, len(tags))
	for i, v := range tags {
		result[i] = v.Value
	}

	return result
}

func (t tags) StringsToTags(userID uint) []*bbs.Tag {
	tags := make([]*bbs.Tag, len(t))
	for i, v := range t {
		tags[i] = &bbs.Tag{UserID: userID, Value: v}
	}
	return tags
}

type TagHandler struct {
	service bbs.TagService
}

func NewTagHandler(s bbs.TagService) TagHandler {
	return TagHandler{s}
}

func (t *TagHandler) ListTag(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	// db.
	tags, err := t.service.Tags()
	if err != nil {
		return nil, NewAppError(err)
	}

	// response.
	strings := newTags(tags)
	data := struct {
		Total int      `json:"total"`
		Tags  []string `json:"tags"`
	}{len(strings), strings}
	return data, nil
}
