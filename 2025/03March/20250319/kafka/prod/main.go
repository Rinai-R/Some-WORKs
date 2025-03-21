package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/IBM/sarama"
)

func main() {
	brokers := []string{"localhost:29092"}
	topic := "test_topic"

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start Kafka async producer: %v", err)
	}
	defer producer.Close()

	// 监听成功和失败的消息
	go func() {
		for msg := range producer.Successes() {
			fmt.Printf("Message sent successfully: %v\n", msg)
		}
	}()

	go func() {
		for err := range producer.Errors() {
			fmt.Printf("Failed to send message: %v\n", err)
		}
	}()
	wg := sync.WaitGroup{}
	wg.Add(1000000)
	// 发送消息
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				msg := &sarama.ProducerMessage{
					Topic: topic,
					Value: sarama.StringEncoder(fmt.Sprintf("Hello Kafka!! My id is %d", i)),
				}
				producer.Input() <- msg
				defer wg.Done()
			}
		}()
	}
	wg.Wait()
}
