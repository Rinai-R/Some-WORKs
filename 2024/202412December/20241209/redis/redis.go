package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// 声明 Redis 客户端的全局变量
var rdb *redis.Client

// 初始化函数
func init() {
	// 创建一个新的 Redis 客户端
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器的地址和端口
		Password: "",               // 没有设置密码
		DB:       0,                // 使用默认数据库
	})
}

func main() {
	// 创建一个上下文
	ctx := context.Background()

	// 将键 "gorediskey" 设置为值 "goredisvalue"，并且设置过期时间为 0，表示永不过期
	err := rdb.Set(ctx, "gorediskey", "goredisvalue", 0).Err()
	if err != nil {
		panic(err) // 如果设置过程中出现错误，终止程序并输出错误
	}

	// 从 Redis 获取键 "gorediskey" 的值
	value, err := rdb.Get(ctx, "text").Result()
	if err != nil {
		panic(err) // 如果获取过程中出现错误，终止程序并输出错误
	}
	fmt.Println("text", value) // 打印键和值

	// 另一种获取方式：使用 Do 方法
	val, err := rdb.Do(ctx, "get", "text").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("text 不存在") // 如果键不存在，输出提示信息
			return
		}
		panic(err) // 如果获取过程中出现其他错误，终止程序并输出错误
	}
	fmt.Println("do operator : text", val.(string)) // 打印通过 Do 获取的值
}
