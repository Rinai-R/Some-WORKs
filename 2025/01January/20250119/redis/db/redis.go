package db

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var Rdb *redis.Client
var err error

func init() {
	dsn := "192.168.195.129:6379"
	Rdb = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "yourpassword",
		DB:       0,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("redis connect success")
}
