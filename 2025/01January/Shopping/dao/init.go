package dao

import "database/sql"

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "root:~Cy710822@tcp(127.0.0.1:3306)/shopping?parseTime=True&loc=UTC")
	if err != nil {
		panic(err)
	}
}
