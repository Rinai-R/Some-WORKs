package main

import (
	"Golang/2025/01January/20250120/redis-test/db"
	"context"
	"fmt"
	"log"
)

var ctx = context.Background()

func main() {
	res, err := db.Rdb.Get(ctx, "shopping:goods:1").Result()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res)
}
