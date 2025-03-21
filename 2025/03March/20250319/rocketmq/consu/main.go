package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"192.168.195.129:9876"}), // 接入点地址
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("ConsumerGroup"), // 分组名称
	)
	c.Subscribe("topicName", consumer.MessageSelector{}, func(ctx context.Context, msg ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, v := range msg {
			fmt.Println(string(v.Body)) // v.Body : 消息主体
		}
		return consumer.ConsumeSuccess, nil
	})
	forever := make(chan bool)
	err := c.Start()
	if err != nil {
		log.Fatal(err)
	}
	<-forever

}
