package data

type Topic struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Intro string `json:"intro"`
	Content string `json:"content"`
}

func Topics() (topics []Topic, err error) {
	rows, err := Db.Query("SELECT id, name, author, intro, content from topic")
	if err != nil {
		return
	}

	for rows.Next() {
		t := Topic{}
		if err = rows.Scan(&t.Id, &t.Name, &t.Author, &t.Intro, &t.Content); err != nil {
			return
		}
		topics = append(topics, t)
	}
	rows.Close()
	return
}

func CreateTopic(name, author, intro, content string) (id int64, err error) {
	statement := "insert into topic (name, author, intro, content) values (?,?,?,?)"
	result, err := Db.Exec(statement, name, author, intro, content)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

func GetTopic(id int64) (topic Topic, err error) {
	err = Db.QueryRow("SELECT id, name, author, content from topic where id = ?", id).
		Scan(&topic.Id, &topic.Name, &topic.Author, &topic.Content)
	return
}