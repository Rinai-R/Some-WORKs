package main

import (
	"Golang/12December/20241210/PurChase/dao"
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(22)
	var ctx context.Context
	var rdb *redis.Client
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "ip",
		Password: "password",
		DB:       0,
	})
	for i := 0; i <= 10; i++ {
		go func() {
			defer wg.Done()
			dao.Producer(ctx, rdb)
		}()
		go func() {
			defer wg.Done()
			dao.Purchase(ctx, rdb)
		}()
	}
	wg.Wait()
}
