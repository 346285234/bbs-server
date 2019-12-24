package router

import (
	"encoding/json"
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Success bool 		`json:"success"`
	Code int 			`json:"code"`
	Message string 		`json:"message"`
	Data interface{} 	`json:"data"`
}


type TopicRouter struct {
	logger log.Logger
}

var Tr = TopicRouter{}

func (_ *TopicRouter)SetupRouter(mux http.ServeMux) {
}

func (_ *TopicRouter)ListTopics(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: analyse request

	// query data
	topics, err := data.Ts.Topics()

	// return response
	var response Response
	if err != nil {
		response = Response{Success: false, Code: 500, Message: "Failed"}
		w.WriteHeader(500)
		http.Error(w, err.Error(), 500)
	} else {
		response = Response{Success: true, Code: 200, Message: "OK"}
		data := struct {
			Total int `json:"total"`
			Topics []data.Topic `json:"topics"`
		}{len(topics), topics}
		response.Data = data
	}

	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (_ *TopicRouter)GetTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		return
	}

	topic, err := data.Ts.GetTopic(uint(id))

	var response Response
	if err != nil {
		response = Response{Success: false, Code: 500, Message: "Failed"}
		w.WriteHeader(500)
		http.Error(w, err.Error(), 500)
	} else {
		response = Response{Success: true, Code: 200, Message: "OK"}
		response.Data = topic
	}

	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}


func (_ *TopicRouter)CreateTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var topic data.Topic
	json.NewDecoder(r.Body).Decode(&topic)
	defer r.Body.Close()
	userID := r.Header.Get("userID")
	topic.AuthorID = uint(common.StrToInt(userID))

	data.Ts.AddTopic(topic)
}
