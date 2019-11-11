package router

import (
	"encoding/json"
	"fmt"
	"github.com/346285234/bbs-server/common"
	"github.com/346285234/bbs-server/data"
	"net/http"
)

func ListTopics(writer http.ResponseWriter, request *http.Request) {
	topics, _ := data.Topics()
	response := struct {
		Total int `json:"total"`
		Topics []data.Topic `json:"topics"`
	}{len(topics), topics}
	bytes, _ := json.Marshal(response)
	writer.Write(bytes)
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

func ReadTopic(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	id := vals.Get("id")
	fmt.Println(id)
}
