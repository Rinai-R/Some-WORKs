package main

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"strconv"
)

func main() {
	prod, _ := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(5),
		producer.WithGroupName("PG"),
	)
	err := prod.Start()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 50000; i++ {
		go func() {
			num := strconv.Itoa(i)
			msg := &primitive.Message{
				Topic: "Num",
				Body:  []byte(num),
			}
			// 发送消息
			_, _ = prod.SendSync(context.Background(), msg)
		}()
	}
	select {}
}
