package dao

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var ctx context.Context
var rdb *redis.Client

func init() {
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "ip",
		Password: "password",
		DB:       0,
	})
}

func Publish(channel string, message string) bool {
	mes := fmt.Sprintf("mess:%v", channel)
	err := rdb.LPush(ctx, mes, message).Err()
	if err != nil {
		log.Println(err)
		return false
	}
	err2 := rdb.Publish(ctx, channel, message).Err()
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

func Subscribe(channel string) []string {
	mes := fmt.Sprintf("mess:%v", channel)
	message, err := rdb.LRange(ctx, mes, 0, -1).Result()
	if err != nil {
		log.Println(err)
		return nil
	}
	for len(message) >= 10 {
		err := rdb.RPop(ctx, mes).Err()
		if err != nil {
			log.Println(err)
			return nil
		}
		message, err = rdb.LRange(ctx, mes, 0, -1).Result()
		if err != nil {
			log.Println(err)
			return nil
		}
	}
	return message
}
