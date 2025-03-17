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
	pipe := rdb.Pipeline()
	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("key-%d", i)
		pipe.Set(ctx, key, "value", 0)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 10000; i++ {
		key := fmt.Sprintf("key-%d", i)
		val, err := rdb.Get(ctx, key).Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(val)
	}
}
