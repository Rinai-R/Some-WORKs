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
	tx := rdb.TxPipeline()
	tx.Set(ctx, "key1", "value1", 0)
	tx.Set(ctx, "key2", "value2", 0)
	tx.Exec(ctx)
	val, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	val, err = rdb.Get(ctx, "key2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
