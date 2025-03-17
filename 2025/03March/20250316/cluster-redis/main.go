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

	if err := rdb.Set(ctx, "name", "value", 0).Err(); err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	err = rdb.ForEachMaster(ctx, func(ctx context.Context, master *redis.Client) error {
		fmt.Println(master, "is master")
		return master.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}
	err = rdb.ForEachSlave(ctx, func(ctx context.Context, slave *redis.Client) error {
		fmt.Println(slave, "is slave")
		return slave.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}
}
