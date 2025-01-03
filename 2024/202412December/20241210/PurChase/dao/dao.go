package dao

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"math/rand"
	"strconv"
)

func Purchase(ctx context.Context, rdb *redis.Client) bool {
	defer rdb.LPush(ctx, "lock", "Get")
	num := rand.Intn(10)
	for {
		err1 := rdb.LPop(ctx, "lock").Err()
		if err1 == redis.Nil {
			continue
		} else if err1 == nil {
			break
		}
	}
	remain, err := rdb.Get(ctx, "goods").Result()
	if err != nil {
		log.Fatal(err)
		return false
	}
	remainNum, _ := strconv.Atoi(remain)
	if remainNum < num {
		fmt.Printf("你想买%d个，现存只有%d个，购买失败，库存不足辣！\n", num, remainNum)
		return false
	}
	err = rdb.DecrBy(ctx, "goods", int64(num)).Err()
	if err != nil {
		log.Fatal(err)
		return false
	}
	fmt.Printf("你买了%d个商品，还剩下%d个\n", num, remainNum-num)
	return true
}

func Producer(ctx context.Context, rdb *redis.Client) bool {
	defer rdb.LPush(ctx, "lock", "Get")
	num := rand.Intn(10)
	for {
		err1 := rdb.LPop(ctx, "lock").Err()
		if err1 == redis.Nil {
			continue
		} else if err1 == nil {
			break
		}
	}

	err := rdb.IncrBy(ctx, "goods", int64(num)).Err()
	if err != nil {
		log.Println("不明错误", err)
		return false
	}
	fmt.Printf("生产了%d个商品\n", num)
	return true
}
