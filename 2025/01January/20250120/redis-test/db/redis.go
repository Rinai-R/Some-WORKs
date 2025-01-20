package db

import (
	"Golang/2025/01January/20250120/redis-test/conf"
	"github.com/redis/go-redis/v9"
	"log"
)

var Rdb *redis.Client

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Conf.Addr,
		Password: conf.Conf.Password,
		DB:       conf.Conf.DB,
		PoolSize: 8,
	})
	log.Println("redis init success")
}
