package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 自定义中间件，记录请求的处理时间
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求开始时间
		start := time.Now()

		// 请求处理前的操作
		log.Println("请求开始")

		// 执行请求处理
		c.Next()

		// 请求处理后的操作
		duration := time.Since(start)
		log.Printf("请求 %s %s 花费了 %v\n", c.Request.Method, c.Request.URL.Path, duration)
	}
}

func main() {
	r := gin.Default()

	// 使用 Logger 中间件
	r.Use(Logger())

	// 定义一个路由
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// 启动 HTTP 服务
	r.Run(":8080")
}
