package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:~Cy710822@tcp(127.0.0.1:3306)/my_new_database")
	if err != nil {
		log.Fatal(err)
	}
}
func SelectUser(username string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"
	err := db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		log.Println("Error querying user:", err)
		return false
	}
	return exists
}

func AddUser(username string, password string) {
	query := "INSERT INTO users(username, password) VALUES(?, ?)"
	_, err := db.Exec(query, username, password)
	if err != nil {
		log.Println("Error inserting user:", err)
		return
	}
}

func SelectPasswordFromUsername(username string) string {
	var password string
	query := "SELECT username, password FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&username, &password)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	return password
}

func AlterPassword(username string, newPassword string) bool {
	query := "UPDATE users SET password = ? WHERE username = ?"
	_, err := db.Exec(query, newPassword, username)
	if err != nil {
		log.Println("Error altering password:", err)
		return false
	}
	return true
}

func DeleteUser(username string) bool {
	query := "DELETE FROM users WHERE username = ?"
	_, err := db.Exec(query, username)
	if err != nil {
		log.Println("Error deleting user:", err)
		return false
	}
	return true
}
