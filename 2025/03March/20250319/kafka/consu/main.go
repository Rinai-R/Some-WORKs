package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:29092"}
	topic := "test_topic"

	// 配置消费者
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// 创建消费者
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
	defer consumer.Close()

	// 选择分区
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to start partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	// 监听消息
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Received message: key=%s, value=%s, partition=%d, offset=%d\n",
			string(msg.Key), string(msg.Value), msg.Partition, msg.Offset)
	}
}
