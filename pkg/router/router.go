package router

import (
	"encoding/json"
	"net/http"

	handler2 "github.com/346285234/bbs-server/pkg/router/handler"
	mux "github.com/julienschmidt/httprouter"
)

type errorHandler func(http.ResponseWriter, *http.Request, mux.Params) (interface{}, *AppError)

type Route struct {
	Method  string
	Path    string
	Handler errorHandler
}

type Routes []Route

var routes = Routes{
	Route{
		Method:  "GET",
		Path:    "/topics",
		Handler: handler2.Th.ListTopic,
	},
	Route{
		Method:  "GET",
		Path:    "/topic/:id",
		Handler: handler2.Th.GetTopic,
	},
	Route{
		Method:  "POST",
		Path:    "/topic/add",
		Handler: checkLogin(handler2.Th.AddTopic),
	},
	Route{
		Method:  "POST",
		Path:    "/topic/remove",
		Handler: checkLogin(handler2.Th.RemoveTopic),
	},
	Route{
		Method:  "POST",
		Path:    "/topic/update",
		Handler: checkLogin(handler2.Th.UpdateTopic),
	},
	Route{
		Method:  "POST",
		Path:    "/favorite/topic/:topic_id/mark",
		Handler: checkLogin(handler2.FaH.MarkFavorite),
	},
	Route{
		Method:  "GET",
		Path:    "/favorite/topic/:topic_id",
		Handler: checkLogin(handler2.FaH.CheckFavorite),
	},
	Route{
		Method:  "POST",
		Path:    "/like/topic/:topic_id/mark",
		Handler: checkLogin(handler2.LiH.MarkLikeTopic),
	},
	Route{
		Method:  "POST",
		Path:    "/like/comment/:comment_id/mark",
		Handler: checkLogin(handler2.LiH.MarkLikeComment),
	},
	Route{
		Method:  "GET",
		Path:    "/like/topic/:topic_id",
		Handler: checkLogin(handler2.LiH.CheckLikeTopic),
	},
	Route{
		Method:  "GET",
		Path:    "/like/comment/:comment_id",
		Handler: checkLogin(handler2.LiH.CheckLikeComment),
	},
	Route{
		Method:  "GET",
		Path:    "/tags",
		Handler: checkLogin(handler2.TaH.ListTag),
	},
	Route{
		Method:  "GET",
		Path:    "/categories",
		Handler: handler2.CaH.ListCategory,
	},
	Route{
		Method:  "GET",
		Path:    "/comments/:topic_id",
		Handler: handler2.Ch.List,
	},
	Route{
		Method:  "POST",
		Path:    "/comment/:topic_id/reply",
		Handler: checkLogin(handler2.Ch.Reply),
	},
	Route{
		Method:  "POST",
		Path:    "/comment/:topic_id/revoke",
		Handler: checkLogin(handler2.Ch.Revoke),
	},
}

func NewRouter() *mux.Router {
	router := mux.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, checkError(route.Handler))
	}

	return router
}

func checkError(fn errorHandler) mux.Handle {
	return func(w http.ResponseWriter, r *http.Request, p mux.Params) {
		var response Response

		data, err := fn(w, r, p)
		if data == nil {
			data = make(map[string]interface{})
		}
		if err != nil {
			response = Response{false, err.Code, err.Message, data}
		} else {
			response = Response{true, 200, "OK", data}
		}

		bytes, _ := json.Marshal(response)
		w.Write(bytes)
	}
}

func checkLogin(fn errorHandler) errorHandler {
	return func(w http.ResponseWriter, r *http.Request, p mux.Params) (interface{}, *AppError) {
		if err := checkUser(); err == nil {
			return fn(w, r, p)
		} else {
			return nil, NewAppError(err)
		}
	}
}

func checkUser() error {
	return nil
}
