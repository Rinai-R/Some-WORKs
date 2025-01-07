package dao

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

var db *sql.DB
var err error
var ctx = context.Background()
var rdb *redis.Client

func init() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/shopping?parseTime=True&loc=UTC")
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.195.128:6379",
		Password: "",
		DB:       0,
	})
}
