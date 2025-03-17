package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			"localhost:7001",
			"localhost:7002",
			"localhost:7003",
			"localhost:7004",
			"localhost:7005",
			"localhost:7006",
		},
	})
	script := `	redis.call("SET", KEYS[1], "Hello World")
				return redis.call("GET", KEYS[1])`
	val, err := rdb.Eval(ctx, script, []string{"key-1"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	defer rdb.Close()
}
