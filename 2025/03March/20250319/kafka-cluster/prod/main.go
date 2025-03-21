package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/IBM/sarama"
	"github.com/fatih/color"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			color.Red("Error: %v", err)
		}
	}()
	brokers := []string{"localhost:29092", "localhost:29093", "localhost:29094"}
	topic := "cluster_test_topic"

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	config.Producer.Idempotent = true
	config.Net.MaxOpenRequests = 1
	// config.Producer.Transaction.ID = "my-transaction-id-11"
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	defer producer.Close()
	// producer.BeginTxn()
	// defer producer.AbortTxn()
	go func() {
		for err := range producer.Errors() {
			panic(err)
		}
	}()
	wg := sync.WaitGroup{}
	n := 10

	go func() {
		for msg := range producer.Successes() {
			fmt.Printf("Message sent to partition %d at offset %d\n", msg.Partition, msg.Offset)
			wg.Done()
		}
	}()
	for i := 0; i < n; i++ {
		msg := &sarama.ProducerMessage{
			Topic:     topic,
			Key:       sarama.StringEncoder(fmt.Sprintf("key-%d", i)),
			Value:     sarama.StringEncoder("testonessssssssssssssssssssssssssssssssssss---"),
			Partition: 0,
		}
		fmt.Println(sarama.StringEncoder("I'm Kafka Encoder!"))
		producer.Input() <- msg
		wg.Add(1)
	}
	// producer.CommitTxn()
	wg.Wait()
}
