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
					Path:    "/favorite/topic/:topic_id",
					Handler: checkLogin(h.CheckFavorite),
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
					Path:    "/like/topic/:topic_id",
					Handler: checkLogin(h.CheckLikeTopic),
				},
				Route{
					Method:  "GET",
					Path:    "/like/comment/:comment_id",
					Handler: checkLogin(h.CheckLikeComment),
				}}
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

	//url := "http://localhost:8201/inner/user/1"
	//resp, err := http.Get(url)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
	//	os.Exit(1)
	//}
	//b, err := ioutil.ReadAll(resp.Body)
	//resp.Body.Close()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	//	os.Exit(1)
	//}
	//fmt.Println("hello world!")
	//fmt.Printf("%s", b)
	//fmt.Println(b)

	return nil
}

type User struct {
}

func getUser(id string) (*User, error) {
	return nil, nil
}
