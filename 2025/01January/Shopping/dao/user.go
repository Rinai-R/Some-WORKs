package dao

import (
	"Golang/2025/01January/Shopping/model"
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// Exist 验证用户名或者该用户是否存在
func Exist(username string) string {
	var Exist bool
	query := `select 1 from user where username = ?`
	err := db.QueryRow(query, username).Scan(&Exist)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
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
		if errors.Is(err, sql.ErrNoRows) {
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

// GetUserInfo 查询用户信息
func GetUserInfo(user *model.User) bool {

	//先查redis
	CacheKey := "user:" + user.Username
	if CacheData, err := rdb.Get(ctx, CacheKey).Result(); err == nil {
		if err := json.Unmarshal([]byte(CacheData), &user); err == nil {
			rdb.Set(ctx, CacheKey, user.Username, time.Hour)
			log.Println("redis缓存读取")
			return true
		}
	}
	//没查到，查mysql
	query := `SELECT id, username, password, balance, avatar, nickname, bio FROM user WHERE username = ?`
	err := db.QueryRow(query, user.Username).Scan(&user.Id, &user.Username, &user.Password, &user.Balance, &user.Avatar, &user.Nickname, &user.Bio)
	if err != nil {
		log.Println(err)
		return false
	}

	CacheData, err0 := json.Marshal(user)
	if err0 != nil {
		log.Println(err0)
		return false
	}
	rdb.Set(ctx, CacheKey, string(CacheData), time.Hour)

	return true
}

// Recharge 充值
func Recharge(money float64, username string) bool {
	query := `UPDATE user SET balance = balance + ? WHERE username = ?`
	_, err := db.Exec(query, money, username)
	if err != nil {
		log.Println(err)
		return false
	}

	CacheKey := "user:" + username
	if CacheData, err := rdb.Get(ctx, CacheKey).Result(); err == nil {
		query = `SELECT balance FROM user WHERE username = ?`
		var balance float64
		err := db.QueryRow(query, username).Scan(&balance)
		if err != nil {
			log.Println(err)
			return false
		}
		var Info model.User
		if err = json.Unmarshal([]byte(CacheData), &Info); err == nil {
			Info.Balance = balance
			cachedata, err := json.Marshal(Info)
			if err != nil {
				log.Println(err)
				return false
			}
			rdb.Set(ctx, CacheKey, cachedata, time.Hour)
		}
	}

	return true
}

// AlterUserInfo 改用户信息
func AlterUserInfo(NewInfo model.User, username string) bool {
	tx, _ := db.Begin()
	if NewInfo.Password != "" {
		query := `UPDATE user SET password = ? WHERE username = ?`
		_, err := tx.Exec(query, NewInfo.Password, username)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return false
			}
			log.Println(err)
			return false
		}
	}

	if NewInfo.Nickname != "" {
		query := `UPDATE user SET nickname = ? WHERE username = ?`
		_, err := tx.Exec(query, NewInfo.Nickname, username)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return false
			}
			log.Println(err)
			return false
		}
	}

	if NewInfo.Avatar != "" {
		query := `UPDATE user SET avatar = ? WHERE username = ?`
		_, err := tx.Exec(query, NewInfo.Avatar, username)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return false
			}
			log.Println(err)
			return false
		}
	}

	if NewInfo.Bio != "" {
		query := `UPDATE user SET bio = ? WHERE username = ?`
		_, err := tx.Exec(query, NewInfo.Bio, username)
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return false
			}
			log.Println(err)
			return false
		}
	}
	//redis缓存库的更新，保证数据一致性
	CacheKey := "user:" + username
	if CacheData, err := rdb.Get(ctx, CacheKey).Result(); err == nil {
		var New model.User
		if json.Unmarshal([]byte(CacheData), &New) == nil {
			if NewInfo.Nickname != "" {
				New.Nickname = NewInfo.Nickname
			}
			if NewInfo.Avatar != "" {
				New.Avatar = NewInfo.Avatar
			}
			if NewInfo.Bio != "" {
				New.Bio = NewInfo.Bio
			}
			if NewInfo.Password != "" {
				New.Password = NewInfo.Password
			}
			cacheData, err := json.Marshal(New)
			if err != nil {
				log.Println(err)
				return false
			}
			rdb.Set(ctx, CacheKey, cacheData, time.Hour)
		}
	}
	return true
}

func DelUser(username string) bool {
	query := `DELETE FROM user WHERE username = ?`
	_, err := db.Exec(query, username)
	if err != nil {
		log.Println(err)
		return false
	}
	CacheKey := "user:" + username
	rdb.Del(ctx, CacheKey)
	return true
}
