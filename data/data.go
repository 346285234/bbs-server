package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:346285234@/bbs")
	if err != nil {
		panic(err)
	}
}