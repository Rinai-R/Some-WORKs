package main

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/fatih/color"
)

// ConsumerGroupHandler 实现了 ConsumerGroupHandler 接口
type ConsumerGroupHandler struct{}

// Cleanup implements sarama.ConsumerGroupHandler.
func (h *ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	color.Green("消费者关闭！\n")
	return nil
}

// Setup implements sarama.ConsumerGroupHandler.
func (h *ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	color.Green("消费者启动！\n")
	return nil
}

// 消费消息并打印
func (h *ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("Consumed message: %s\n", string(message.Value))
		// 手动提交偏移量
		sess.MarkMessage(message, "")
	}
	return nil
}

var _ sarama.ConsumerGroupHandler = (*ConsumerGroupHandler)(nil)

func main() {
	defer func() {
		if err := recover(); err != nil {
			color.Red("Error: %v", err)
		}
	}()
	brokers := []string{"localhost:29092", "localhost:29093", "localhost:29094"}
	topic := "cluster_test_topic"

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	ConsumerGroup, err := sarama.NewConsumerGroup(brokers, "my-group", config)
	if err != nil {
		panic(err)
	}
	ConsumerGroup.
	go func() {
		for {
			err = ConsumerGroup.Consume(context.Background(), []string{topic}, &ConsumerGroupHandler{})
			if err != nil {
				color.Red("Error from consumer: %v", err)
			}
		}
	}()
	defer ConsumerGroup.Close()
	select {}
}
