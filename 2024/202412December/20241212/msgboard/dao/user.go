package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func Register(username string, nickname string, password string) string {
	var Exist bool
	query1 := "SELECT 1 FROM users WHERE username = ? "
	db.QueryRow(query1, username).Scan(&Exist)
	if Exist {
		return "exist"
	}
	query := "INSERT INTO users(username, nickname, password) VALUES(?,?,?)"
	idGet, err := db.Exec(query, username, nickname, password)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	id, err2 := idGet.LastInsertId()
	if err2 != nil {
		log.Fatal(err)
	}
	ID := strconv.FormatInt(id, 10)
	return ID
}

func Login(username string, password string) bool {
	query1 := "SELECT 1 FROM users WHERE username = ? "
	var exist bool
	db.QueryRow(query1, username).Scan(&exist)
	if exist == false {
		return false
	}
	query := "SELECT password FROM users WHERE username= ? "
	var TruePassword string
	err := db.QueryRow(query, username).Scan(&TruePassword)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if TruePassword != password {
		return false
	}
	return true
}

func DeleteUser(username string) bool {
	query := "DELETE FROM users WHERE username = ?"
	_, err := db.Exec(query, username)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func GetUserId(username string) string {
	query := "SELECT id FROM users WHERE username=?"
	var UserID string
	err := db.QueryRow(query, username).Scan(&UserID)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return UserID
}
