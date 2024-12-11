package dao

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "ip",
		Password: "~password",
		DB:       0,
	})
}

func RegisterUser(name string, password string, email string, sex string) string {
	userID, err := rdb.Incr(ctx, "user:next_id").Result()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	Userid := fmt.Sprintf("%d", userID)
	err1 := rdb.HSet(ctx, Userid, "name", name, "password", password, "email", email, "sex", sex).Err()
	if err1 != nil {
		log.Fatal(err1)
		return ""
	}
	err2 := rdb.Expire(ctx, Userid, 30*time.Second).Err()
	if err2 != nil {
		log.Fatal(err2)
		return ""
	}
	return Userid
}

func Login(UserID string, password string) bool {
	defer rdb.LPush(ctx, "lock", "Get")
	for {
		err1 := rdb.LPop(ctx, "lock").Err()
		if err1 == redis.Nil {
			continue
		} else if err1 == nil {
			break
		}
	}
	TruePassword, err1 := rdb.HGet(ctx, UserID, "password").Result()
	if err1 != nil {
		log.Fatal(err1)
		return false
	}
	if TruePassword != password {
		return false
	}
	err2 := rdb.Expire(ctx, UserID, 30*time.Second).Err()
	if err2 != nil {
		log.Fatal(err2)
		return false
	}
	return true
}

func AlterPassword(UserID string, password string, NewPassword string) bool {
	TruePassword, err1 := rdb.HGet(ctx, UserID, "password").Result()
	if err1 != nil {
		log.Fatal(err1)
		return false
	}
	if TruePassword != password {
		return false
	}
	err2 := rdb.HSet(ctx, UserID, "password", NewPassword).Err()
	if err2 != nil {
		log.Fatal(err2)
		return false
	}
	err3 := rdb.Expire(ctx, UserID, 30*time.Second).Err()
	if err3 != nil {
		log.Fatal(err3)
		return false
	}
	return true
}

func GetUser(UserID string, password string) (string, string, string) {
	if !Login(UserID, password) {
		return "", "", ""
	}
	name, err1 := rdb.HGet(ctx, UserID, "name").Result()
	if err1 != nil {
		log.Fatal(err1)
		return "", "", ""
	}
	sex, err2 := rdb.HGet(ctx, UserID, "sex").Result()
	if err2 != nil {
		log.Fatal(err2)
		return "", "", ""
	}
	email, err3 := rdb.HGet(ctx, UserID, "email").Result()
	if err3 != nil {
		log.Fatal(err3)
		return "", "", ""
	}
	return name, sex, email
}
