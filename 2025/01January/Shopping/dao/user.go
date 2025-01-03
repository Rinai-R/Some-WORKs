package dao

import (
	"Golang/2025/01January/Shopping/model"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// Exist 验证用户名或者该用户是否存在
func Exist(username string) string {
	var Exist bool
	query := `select 1 from user where username = ?`
	err := db.QueryRow(query, username).Scan(&Exist)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return "none"
		}
		log.Println(err)
		return "error"
	}
	if Exist {
		return "exists"
	}
	return "none"
}

func Register(user model.User) bool {
	query := `insert into user (username, password) values (?, ?)`
	_, err1 := db.Exec(query, user.Username, user.Password)
	if err1 != nil {
		log.Println(err1)
		return false
	}
	id := GetId(user.Username)
	if id == "" {
		return false
	}
	query = `insert into shopping_cart (user_id) values (?)`
	_, err3 := db.Exec(query, id)
	if err3 != nil {
		log.Println(err3)
		return false
	}
	return true
}

func Login(user model.User) string {
	query := `select 1 from user where username = ? and password = ?`
	var exist bool
	err := db.QueryRow(query, user.Username, user.Password).Scan(&exist)
	if err != nil {
		if err == sql.ErrNoRows {
			return "wrong password or username"
		}
		log.Println(err)
		return "error"
	}
	if exist {
		return "ok"
	}
	return "wrong password or username"
}

func GetId(username string) string {
	query := `select id from user where username = ?`
	var id string
	err := db.QueryRow(query, username).Scan(&id)
	if err != nil {
		log.Println(err)
		return ""
	}
	return id
}

func GetUserInfo(user *model.User) bool {
	query := `SELECT id, username, password, balance, avatar FROM user WHERE username = ?`
	err := db.QueryRow(query, user.Username).Scan(&user.Id, &user.Username, &user.Password, &user.Balance, &user.Avatar)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func Recharge(user model.User) bool {
	query := `UPDATE user SET balance = balance + ? WHERE username = ?`
	_, err := db.Exec(query, user.Balance, user.Username)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func AlterUserInfo(NewInfo model.User, id string) bool {
	if NewInfo.Password != "" {
		query := `UPDATE user SET password = ? WHERE username = ?`
		_, err := db.Exec(query, NewInfo.Password, id)
		if err != nil {
			log.Println(err)
			return false
		}
	}

	if NewInfo.Username != "" {
		if Exist(NewInfo.Username) != "none" {
			return false
		}
		query := `UPDATE user SET username = ? WHERE id = ?`
		_, err := db.Exec(query, NewInfo.Username, id)
		if err != nil {
			log.Println(err)
			return false
		}
	}

	if NewInfo.Avatar != "" {
		query := `UPDATE user SET avatar = ? WHERE id = ?`
		_, err := db.Exec(query, NewInfo.Avatar, id)
		if err != nil {
			log.Println(err)
			return false
		}
	}
	return true
}
