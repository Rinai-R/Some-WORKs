package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/segmentio/kafka-go"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			color.Red("panic: %v", err)
		}
	}()
	writer := kafka.Writer{
		Addr:         kafka.TCP("localhost:29092", "localhost:29093", "localhost:29094"),
		Topic:        "cluster_test_topic",
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: 10 * time.Millisecond,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 发送消息
	err := writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte("key-1"),
		Value: []byte("Hello World Kafka!"),
	})
	if err != nil {
		panic("发送消息失败: " + err.Error())
	}

	fmt.Println("消息发送成功！")
}
