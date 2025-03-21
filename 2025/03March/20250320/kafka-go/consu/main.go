package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:          []string{"localhost:29092", "localhost:29093", "localhost:29094"},
		Topic:            "cluster_test_topic",
		MinBytes:         1,
		MaxBytes:         1e6,
		GroupID:          "mygroup",
		ReadBatchTimeout: 10 * time.Millisecond,
	})
	defer reader.Close()
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("message at offset %d: %s\n", m.Offset, string(m.Value))
	}

}
