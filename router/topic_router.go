package router

import (
	"encoding/json"
	"fmt"
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data"
	"github.com/346285234/bbs-server/data/services"
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

var Tr = TopicRouter{}

func (_ *TopicRouter)SetupRouter(mux http.ServeMux) {
}

func (_ *TopicRouter)ListTopics(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: analyse request

	// query data
	topics, err := services.Ts.Topics()

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
		}{len(*topics), *topics}
		response.Data = data
	}

	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

func (_ *TopicRouter)GetTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var id uint = 1 // p.ByName("id").(uint)

	// query data
	topics, err := services.Ts.GetTopic(id)

	// return response
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
	//	}{len(*topics), *topics}
	//	response.Data = data
	//}
	//
	//bytes, _ := json.Marshal(response)
	//w.Write(bytes)
}

//
//func CreateTopic(writer http.ResponseWriter, request *http.Request) {
//	var response struct {
//		Para data.Topic `json:"parameters"`
//	}
//	json.NewDecoder(request.Body).Decode(&response)
//	defer request.Body.Close()
//	introV := common.Intro(response.Para.Content)
//	id, _ := data.CreateTopic(response.Para.Name, response.Para.Author, introV, response.Para.Content)
//	fmt.Println(id)
//}
