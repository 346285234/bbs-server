// Package handler implements all http handler.
package router

import (
	"github.com/346285234/bbs-server/pkg/bbs"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CategoriesResponse []CategoryResponse

func NewCategoriesResponse(categories []bbs.Category) CategoriesResponse {
	result := make([]CategoryResponse, len(categories))
	for i, v := range categories {
		result[i] = CategoryResponse{v.ID, v.Value}
	}
	return result
}

type CategoryHandler struct {
	service bbs.CategoryService
}

func NewCategoryHandler(s bbs.CategoryService) CategoryHandler {
	return CategoryHandler{s}
}

func (c *CategoryHandler) ListCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, *AppError) {
	// db.
	categories, err := c.service.Categories()
	if err != nil {
		return nil, NewAppError(err)
	}

	// response.
	response := NewCategoriesResponse(categories)
	data := struct {
		Total      int                `json:"total"`
		Categories []CategoryResponse `json:"categories"`
	}{len(response), response}
	return data, nil
}
