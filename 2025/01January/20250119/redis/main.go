package main

import (
	"Golang/2025/01January/20250119/redis/db"
	"context"
	"fmt"
	"time"
)

func main() {
	db.Rdb.Set(context.Background(), "hello", "world", time.Minute)
	fmt.Println("hello world")
}
