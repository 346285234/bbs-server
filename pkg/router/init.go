package router

import (
	"encoding/json"
	mux "github.com/julienschmidt/httprouter"
	"net/http"
)

type errorHandler func(http.ResponseWriter, *http.Request, mux.Params) (interface{}, *AppError)

type Route struct {
	Method  string
	Path    string
	Handler errorHandler
}

type Routes []Route

func NewRouter(handlers []interface{}) *mux.Router {
	var routes Routes

	for _, v := range handlers {
		var rs Routes
		switch h := v.(type) {
		case TopicHandler:
			rs = Routes{
				Route{
					Method:  "GET",
					Path:    "/topics",
					Handler: h.ListTopic,
				},
				Route{
					Method:  "GET",
					Path:    "/topic/:topic_id",
					Handler: h.GetTopic,
				},
				Route{
					Method:  "POST",
					Path:    "/topic/add",
					Handler: checkLogin(h.AddTopic),
				},
				Route{
					Method:  "POST",
					Path:    "/topic/remove",
					Handler: checkLogin(h.RemoveTopic),
				},
				Route{
					Method:  "POST",
					Path:    "/topic/update",
					Handler: checkLogin(h.UpdateTopic),
				}}
		case FavoriteHanlder:
			rs = Routes{
				Route{
					Method:  "POST",
					Path:    "/favorite/topic/:topic_id/mark",
					Handler: checkLogin(h.MarkFavorite),
				},
				Route{
					Method:  "GET",
					Path:    "/favorite/topic/:topic_id/check",
					Handler: checkLogin(h.CheckFavorite),
				},
				Route{
					Method:  "GET",
					Path:    "/favorite/topic/:topic_id",
					Handler: h.FavoriteUsers,
				},
			}
		case LikeHandler:
			rs = Routes{
				Route{
					Method:  "POST",
					Path:    "/like/topic/:topic_id/mark",
					Handler: checkLogin(h.MarkLikeTopic),
				},
				Route{
					Method:  "POST",
					Path:    "/like/comment/:comment_id/mark",
					Handler: checkLogin(h.MarkLikeComment),
				},
				Route{
					Method:  "GET",
					Path:    "/like/topic/:topic_id/check",
					Handler: checkLogin(h.CheckLikeTopic),
				},
				Route{
					Method:  "GET",
					Path:    "/like/comment/:comment_id/check",
					Handler: checkLogin(h.CheckLikeComment),
				},
				Route{
					Method:  "GET",
					Path:    "/like/topic/:topic_id",
					Handler: h.likeTopicUsers,
				},
				Route{
					Method:  "GET",
					Path:    "/like/comment/:comment_id",
					Handler: h.likeCommentUsers,
				},
			}
		case CategoryHandler:
			rs = Routes{
				Route{
					Method:  "GET",
					Path:    "/categories",
					Handler: h.ListCategory,
				},
			}
		case TagHandler:
			rs = Routes{
				Route{
					Method:  "GET",
					Path:    "/tags",
					Handler: checkLogin(h.ListTag),
				}}
		case CommentHandler:
			rs = Routes{
				Route{
					Method:  "GET",
					Path:    "/comments/:topic_id",
					Handler: h.List,
				},
				Route{
					Method:  "POST",
					Path:    "/comment/:topic_id/reply",
					Handler: checkLogin(h.Reply),
				},
				Route{
					Method:  "POST",
					Path:    "/comment/:topic_id/revoke",
					Handler: checkLogin(h.Revoke),
				}}
		case FollowHanlder:
			rs = Routes{
				Route{
					Method:  "POST",
					Path:    "/follow/user/:user_id/mark",
					Handler: checkLogin(h.Mark),
				},
				Route{
					Method:  "GET",
					Path:    "/follow/user/:user_id/check",
					Handler: checkLogin(h.Check),
				},
				Route{
					Method:  "GET",
					Path:    "/follow/user/:user_id",
					Handler: h.FollowUsers,
				},
			}
		}

		routes = append(routes, rs...)
	}

	router := mux.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, checkError(route.Handler))
	}

	return router
}

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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
		w.Header().Set("content-type", "application/json")
		w.Write(bytes)

	}
}

func checkLogin(fn errorHandler) errorHandler {
	return func(w http.ResponseWriter, r *http.Request, p mux.Params) (interface{}, *AppError) {
		userID := r.Header.Get("userID")
		if err := checkUser(userID); err == nil {
			return fn(w, r, p)
		} else {
			return nil, NewAppError(err)
		}
	}
}

func checkUser(id string) error {

	return nil
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Portrait string `json:"portrait"`
}
