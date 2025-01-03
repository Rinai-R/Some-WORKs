package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:~Cy710822@tcp(127.0.0.1:3306)/employees")
	if err != nil {
		log.Fatal(err)
	}
}

func Search(id string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM employees WHERE id = ?)"
	err := db.QueryRow(query, id).Scan(&exists)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return exists
}

func AddUser(id string, name string, email string, phone string, hire_date time.Time, salary string, department string, manager_id string, status string, created_at time.Time, updated_at time.Time) {
	query := "INSERT INTO employees(id, name, email, phone, hire_date, salary, department, manager_id, status, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ? ,? ,? ,? ,? ,?)"
	_, err := db.Exec(query, id, name, email, phone, hire_date, salary, department, manager_id, status, created_at, updated_at)
	if err != nil {
		log.Println("Error inserting user:", err)
		return
	}
}
