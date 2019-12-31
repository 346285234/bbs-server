package router

import (
	"encoding/json"
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
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

var tr = TopicRouter{}

func (_ *TopicRouter)SetupRouter(mux http.ServeMux) {
}

func (_ *TopicRouter)getTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) *appError {

	// Get id.
		id := common.StrToInt(p.ByName("id"))

		// Get data.
		topic, err := data.Ts.GetTopic(uint(id))

	if err != nil {
		return &appError{err, "", 500}
	}

		// Set response.
		var response Response
		if err != nil {
			http.Error(w, err.Error(), 500)
			response = Response{Success: false, Code: 500, Message: "Failed"}
		} else {
			response = Response{Success: true, Code: 200, Message: "OK"}
			response.Data = topic
		}

		bytes, _ := json.Marshal(response)
		w.Write(bytes)
}

//func (_ *TopicRouter)listTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	// TODO: Analyse query.
//
//	// Get data.
//	topics, err := data.Ts.Topics()
//
//	// Set response.
//	var response Response
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		response = Response{Success: false, Code: 500, Message: "Failed"}
//	} else {
//		response = Response{Success: true, Code: 200, Message: "OK"}
//		data := struct {
//			Total int `json:"total"`
//			Topics []data.Topic `json:"topics"`
//		}{len(topics), topics}
//		response.Data = data
//	}
//	bytes, _ := json.Marshal(response)
//	w.Write(bytes)
//}
//
//func (_ *TopicRouter)getTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//
//	// Get id.
//	id := common.StrToInt(p.ByName("id"))
//
//	// Get data.
//	topic, err := data.Ts.GetTopic(uint(id))
//
//	// Set response.
//	var response Response
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		response = Response{Success: false, Code: 500, Message: "Failed"}
//	} else {
//		response = Response{Success: true, Code: 200, Message: "OK"}
//		response.Data = topic
//	}
//
//	bytes, _ := json.Marshal(response)
//	w.Write(bytes)
//}
//
//
//func (_ *TopicRouter)addTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	// Analyse response.
//	var topic data.Topic
//	json.NewDecoder(r.Body).Decode(&topic)
//	defer r.Body.Close()
//	userID := r.Header.Get("userID")
//	topic.AuthorID = uint(common.StrToInt(userID))
//
//	// Add data.
//	err := data.Ts.AddTopic(topic)
//
//	// Set response.
//	var response Response
//	if err != nil {
//		response = Response{Success: false, Code: 500, Message: "Failed"}
//		http.Error(w, err.Error(), 500)
//	} else {
//		response = Response{Success: true, Code: 200, Message: "OK"}
//	}
//
//	bytes, _ := json.Marshal(response)
//	w.Write(bytes)
//
//}
//
//func (_ *TopicRouter)removeTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	// Get id.
//	var body map[string]int
//	json.NewDecoder(r.Body).Decode(&body)
//	defer r.Body.Close()
//	id, ok := body["id"]
//	if !ok {
//		// TODO: handle error.
//	}
//
//	topicID := uint(id)
//	userID := uint(common.StrToInt(r.Header.Get("userID")))
//
//	err := data.Ts.RemoveTopic(userID, topicID)
//
//	var response Response
//	if err != nil {
//		response = Response{Success: false, Code: 500, Message: "Failed"}
//		http.Error(w, err.Error(), 500)
//	} else {
//		response = Response{Success: true, Code: 200, Message: "OK"}
//	}
//
//	bytes, _ := json.Marshal(response)
//	w.Write(bytes)
//
//}
//
//func (_ *TopicRouter)updateTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	var topic data.Topic
//	json.NewDecoder(r.Body).Decode(&topic)
//	defer r.Body.Close()
//	userID := r.Header.Get("userID")
//	topic.AuthorID = uint(common.StrToInt(userID))
//
//	err := data.Ts.AddTopic(topic)
//
//	var response Response
//	if err != nil {
//		response = Response{Success: false, Code: 500, Message: "Failed"}
//		http.Error(w, err.Error(), 500)
//	} else {
//		response = Response{Success: true, Code: 200, Message: "OK"}
//	}
//
//	bytes, _ := json.Marshal(response)
//	w.Write(bytes)
//
//}