package router

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type RouterTopic struct {
	logger log.Logger
}

var Rt = RouterTopic{}

func (_ *RouterTopic)SetupRouter(mux http.ServeMux) {
}

func (_ *RouterTopic)ListTopics(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//topics, err := data.Topics()
	//var response Response
	//if err != nil {
	//	response = Response{Success: false, Code: 500, Message: "Failed"}
	//	w.WriteHeader(500)
	//	http.Error(w, err.Error(), 500)
	//} else {
	//	response = Response{Success: true, Code: 200, Message: "OK"}
	//	data := struct {
	//		Total int `json:"total"`
	//		Topics []data.Topic `json:"topics"`
	//	}{len(topics), topics}
	//	response.Data = data
	//}
	//
	//bytes, _ := json.Marshal(response)
	//w.Write(bytes)
}