package data

//type Post struct {
//	Id int
//	Content string
//	Topic_Id int
//}
//
//func (topic *Topic) Posts() (posts []Post, err error) {
//	rows, err := Db.Query("select id, content, topic_id from post where topic_id = ?", topic.Id)
//	if err != nil {
//		return
//	}
//
//	for rows.Next() {
//		post := Post{}
//		err = rows.Scan(&post.Id, &post.Content, &post.Topic_Id)
//		if err != nil {
//			return
//		}
//		posts = append(posts, post)
//	}
//	rows.Close()
//	return
//}
//
//func (topic *Topic) NumberOfPosts() (count int) {
//	rows, err := Db.Query("select count(*) from post where topic_id = ?", topic.Id)
//		if err != nil {
//		return
//	}
//
//	for rows.Next() {
//		if err = rows.Scan(&count); err != nil {
//			return
//		}
//	}
//	rows.Close()
//	return
//}
//
//func (topic *Topic) CreatePost(content string) (post Post, err error) {
//	statement := "insert into post (content, topic_id) values (?, ?)"
//	result, err := Db.Exec(statement, content, topic.Id)
//	if err != nil {
//		return
//	}
//	id, err := result.LastInsertId()
//	post = Post{int(id), content, topic.Id}
//	return
//}