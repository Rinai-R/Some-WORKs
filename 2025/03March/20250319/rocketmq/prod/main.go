package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	p, _ := rocketmq.NewProducer(
		producer.WithNameServer([]string{"192.168.195.129:9876"}), // 接入点地址
		producer.WithRetry(2),                  // 重试次数
		producer.WithGroupName("ProductGroup"), // 分组名称
	)
	err := p.Start()
	if err != nil {
		return
	}
	// 发送的消息
	for i := 0; i < 10; i++ {
		go func() {
			msg := &primitive.Message{
				Topic: "topicName",
				Body:  []byte("Hello RocketMQ Go Client!"),
			}
			msg.WithTag("my-tag")
			msg.WithKeys([]string{"my-key"})
			// 发送消息
			result, _ := p.SendSync(context.Background(), msg)
			fmt.Println(result)
		}()
	}

	select {}

}
