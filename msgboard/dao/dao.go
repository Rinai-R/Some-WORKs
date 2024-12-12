package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/msg?parseTime=True&loc=UTC")
	if err != nil {
		log.Fatal(err)
		return
	}
	//if !ExecEvent() {
	//	return
	//}
}

//func ExecEvent() bool {
//	query := "SET GLOBAL event_scheduler = ON"
//	_, err := db.Exec(query)
//	if err != nil {
//		log.Fatal(err)
//		return false
//	}
//	query2 := "ALTER EVENT delete_expired_data ENABLE"
//	_, err = db.Exec(query2)
//	if err != nil {
//		log.Fatal(err)
//		return false
//	}
//	return true
//}
