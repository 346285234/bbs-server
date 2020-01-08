package router

import (
	"encoding/json"
	mux "github.com/julienschmidt/httprouter"
	"net/http"
)

type appError struct {
	error   error
	Message string
	Code    int
}

func (ae *appError) Error() string {
	return ae.Message + string(ae.Code)
}

func NewAppError(e error) *appError {
	return &appError{e, "", 500}
}

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type errorHandler func(http.ResponseWriter, *http.Request, mux.Params) *appError

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
		Handler: tr.listTopic,
	},
	Route{
		Method:  "GET",
		Path:    "/topic/:id",
		Handler: tr.getTopic,
	},
	Route{
		Method:  "POST",
		Path:    "/topic/add",
		Handler: checkLogin(tr.addTopic),
	},
	Route{
		Method:  "POST",
		Path:    "/topic/remove",
		Handler: checkLogin(tr.removeTopic),
	},
	Route{
		Method:  "POST",
		Path:    "/topic/update",
		Handler: checkLogin(tr.updateTopic),
	},
	//Route{
	//	Method:  "POST",
	//	Path:    "/topic/favourites/mark",
	//	Handler: Tr.MarkFavourites,
	//},
	//Route{
	//	Method:  "POST",
	//	Path:    "/topic/like/mark",
	//	Handler: Tr.markLike,
	//},
	Route{
		Method:  "GET",
		Path:    "/tags",
		Handler: tr.listTag,
	},
	Route{
		Method:  "GET",
		Path:    "/categories",
		Handler: tr.listCategory,
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
		if err := fn(w, r, p); err != nil {
			res := Response{Success: false, Code: err.Code, Message: err.Message}
			bytes, _ := json.Marshal(res)
			w.WriteHeader(err.Code)
			w.Write(bytes)
		}
	}
}

func checkLogin(fn func(w http.ResponseWriter, r *http.Request, p mux.Params) *appError) errorHandler {
	return func(w http.ResponseWriter, r *http.Request, p mux.Params) *appError {
		if err := checkUser(); err == nil {
			return fn(w, r, p)
		} else {
			return &appError{err, "not login", 500}
		}
	}
}

func checkUser() error {
	return nil
}
