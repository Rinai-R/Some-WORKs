package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.195.128:6379",
		Password: "~Cy710822",
		DB:       0,
	})
	UserID := "114514"
	err := rdb.HSet(ctx, UserID, "username", "rina", "email", "lanshan@114514.com").Err()
	if err != nil {
		fmt.Println(err)
	}
}
