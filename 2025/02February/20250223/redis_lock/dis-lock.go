package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         "localhost:7001",
		DB:           0,
		PoolSize:     500, // 默认是10×CPU核心数
		MinIdleConns: 50,  // 保持最小空闲连接
	})
}

type DistributedLock struct {
	client     *redis.Client
	key        string
	identifier string
	stopChan   chan struct{}
}

// 获取锁，expiration为锁的过期时间
func AcquireLock(client *redis.Client, key string, expiration time.Duration) (*DistributedLock, error) {
	identifier := uuid.New().String()

	for {
		// 原子化设置锁（SET NX PX）
		ok, err := client.SetNX(context.Background(), key, identifier, expiration).Result()
		if err != nil {
			return nil, err
		}
		if ok {
			lock := &DistributedLock{
				client:     client,
				key:        key,
				identifier: identifier,
				stopChan:   make(chan struct{}),
			}
			// 启动自动续期
			go lock.autoRefresh(expiration)
			return lock, nil
		}
		m := rand.Intn(1000) + 1000
		time.Sleep(time.Microsecond * time.Duration(m))
	}
}

// 自动续期
func (dl *DistributedLock) autoRefresh(expiration time.Duration) {
	ticker := time.NewTicker(expiration / 3)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 使用Lua脚本保证原子续期
			script := `
			if redis.call("GET", KEYS[1]) == ARGV[1] then
				return redis.call("PEXPIRE", KEYS[1], ARGV[2])
			end
			return 0
			`
			_, err := dl.client.Eval(
				context.Background(),
				script,
				[]string{dl.key},
				dl.identifier,
				expiration.Milliseconds(),
			).Result()

			if err != nil {
				return // 续期失败，退出自动续期
			}
		case <-dl.stopChan:
			return // 收到解锁信号
		}
	}
}

// 释放锁
func (dl *DistributedLock) Release() error {
	close(dl.stopChan) // 停止自动续期

	script := `
	local val = redis.call('GETDEL', KEYS[1])
	if val == ARGV[1] then return 1 end
	return 0
	`
	_, err := dl.client.Eval(
		context.Background(),
		script,
		[]string{dl.key},
		dl.identifier,
	).Result()
	return err
}

var x = 0

func Add() {
	lock, _ := AcquireLock(rdb, "x:lock:add", time.Millisecond)
	x += 1
	fmt.Println("x:", x)
	err := lock.Release()
	if err != nil {
		return
	}
	defer wg.Done()
}

var wg sync.WaitGroup

func main() {
	now := time.Now()
	n := 10000
	wg.Add(n)
	for i := 0; i < n; i++ {
		go Add()
	}
	wg.Wait()
	fmt.Println(time.Since(now))
}
