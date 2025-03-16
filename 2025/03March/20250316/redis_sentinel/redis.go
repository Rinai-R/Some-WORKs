package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// func main() {
// 	sentinel := redis.NewSentinelClient(
// 		&redis.Options{
// 			Addr: ":26379", // 替换为其中一个哨兵的地址

// 		},
// 	)
// 	ctx := context.Background()
// 	addr, err := sentinel.GetMasterAddrByName(ctx, "mymaster").Result()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Master address:", addr)

// 	sentinel.Close()
// }

func main() {
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{":26379", ":26380", ":26381"},
	})
	fmt.Println("Redis client created")
	defer client.Close()
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping result:", pong)
}
