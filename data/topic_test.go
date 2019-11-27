package data
//
//import (
//	"fmt"
//	"testing"
//)
//
//func init() {
//	//deleteTopic()
//}
//
//func deleteTopic() (err error) {
//	fmt.Println("clean")
//	statement := "delete from topic"
//	_, err = Db.Exec(statement)
//
//	return
//}
//
//func TestCreateTopic(t *testing.T) {
//	name := "f"
//	author := "cq"
//	intro := "first"
//	content := "first topic"
//	id, err := CreateTopic(name, author, intro, content)
//	if err != nil {
//		t.Error(err, "Cannot create topic")
//	}
//	fmt.Println(id)
//}
//
//func TestTopics(t *testing.T) {
//	topics, err := Topics()
//	if err != nil {
//		t.Error(err, "Cannot list topics")
//	}
//	fmt.Println(topics)
//}
//
//func TestGetTopic(t *testing.T) {
//	topic, err := GetTopic(1)
//	if err != nil {
//		t.Error(err, "Cannot get topic")
//	}
//	fmt.Println(topic)
//}