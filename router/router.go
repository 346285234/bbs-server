package router

import (
	mux "github.com/julienschmidt/httprouter"
	"net/http"
)

type appError struct {
	Error error
	Message string
	Code int
}

type appHandler func(http.ResponseWriter, *http.Request, mux.Params) *appError
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, p mux.Params) {
	if err := fn(w, r, p); err != nil {
	}
}

type Route struct {
	Method string
	Path string
	Handler mux.Handle
}

type Routes []Route

var routes = Routes{
	//Route{
	//	Method:  "GET",
	//	Path:    "/topics",
	//	Handler: appHandler(tr.listTopic),
	//},
	Route{
		Method:  "GET",
		Path:    "/topic/:id",
		Handler: appHandler(tr.getTopic),
	},
	//Route{
	//	Method:  "POST",
	//	Path:    "/topic/add",
	//	Handler: tr.addTopic,
	//},
	//Route{
	//	Method:  "POST",
	//	Path:    "/topic/remove",
	//	Handler: tr.removeTopic,
	//},
	//Route{
	//	Method:  "POST",
	//	Path:    "/topic/update",
	//	Handler: tr.updateTopic,
	//},
	//Route{
	//	Method:  "POST",
	//	Path:    "topic/favourites/mark",
	//	Handler: Tr.MarkFavourites,
	//},
	//Route{
	//	Method:  "POST",
	//	Path:    "topic/like/mark",
	//	Handler: Tr.markLike,
	//},
	//Route{
	//	Method:  "GET",
	//	Path:    "topic/tags",
	//	Handler: Tr.listTag,
	//},
	//Route{
	//	Method:  "GET",
	//	Path:    "topic/categories",
	//	Handler: Tr.listCategory,
	//},
}

func NewRouter() *mux.Router {
	router := mux.New()
	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.Handler)
	}

	return router
}

//func check(fn func (w http.ResponseWriter,
//	r *http.Request,
//	p httprouter.Params)) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//		if !checkUser() {
//			return
//		}
//		fn(w, r, p)
//	}
//}
//
//func checkUser() bool {
//	// TODO: Check user info from header
//	return true
//}