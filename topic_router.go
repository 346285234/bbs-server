package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data"
	"github.com/julienschmidt/httprouter"
)

type ResponseHandler func(w http.ResponseWriter, r *http.Request) (interface{}, error)

func (rh ResponseHandler)responseHandle(w http.ResponseWriter, r *http.Request) {
	var response Response
	if result, err := rh(w, r); err != nil {
		response = Response{Success: false, Code: 500, Message: "Failed"}
	} else {
		response = Response{Success: true, Code: 200, Message: "OK", Data: result}
	}
	bytes, _ := json.Marshal(response)
	w.Write(bytes)
}

type Response struct {
	Success bool 		`json:"success"`
	Code int 			`json:"code"`
	Message string 		`json:"message"`
	Data interface{} 	`json:"data"`
}

func ListTopics(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	topics, err := data.Topics()
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

func GetTopic(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, p.ByName("id"))
}

func CreateTopic(writer http.ResponseWriter, request *http.Request) {
	var response struct {
		Para data.Topic `json:"parameters"`
	}
	json.NewDecoder(request.Body).Decode(&response)
	defer request.Body.Close()
	introV := common.Intro(response.Para.Content)
	id, _ := data.CreateTopic(response.Para.Name, response.Para.Author, introV, response.Para.Content)
	fmt.Println(id)
}

