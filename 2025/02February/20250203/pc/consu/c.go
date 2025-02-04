package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"log"
	"strconv"
	"sync"
)

var lock sync.Mutex

func main() {
	var sum int
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"192.168.195.129:9876"}), // 接入点地址
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("CG"), // 分组名称
	)
	if err != nil {
		log.Fatal(err)
	}
	forever := make(chan bool)

	// 创建一个通道用来接收消息体
	for i := 0; i < 150; i++ {
		messages := make(chan string)

		// 订阅主题，并定义消息处理函数
		err = c.Subscribe("Num", consumer.MessageSelector{}, func(ctx context.Context, msg ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			for _, v := range msg {
				fmt.Println(string(v.Body)) // v.Body : 消息主体
				messages <- string(v.Body)  // 将消息主体发送到messages通道
			}
			return consumer.ConsumeSuccess, nil
		})
		if err != nil {
			log.Fatal(err)
		}

		// 启动消费者
		err = c.Start()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Consumer started, waiting for messages...")

		go func() {
			for {
				select {
				case msg := <-messages:
					_, err := strconv.Atoi(msg)
					if err != nil {
						log.Println("Error converting message to int:", err)
						continue
					}
					sum += 1
					fmt.Println("Current sum:", sum) // 打印当前的 sum 值
				case <-forever:
					return
				}
			}
		}()
	}

	// 阻塞main函数，防止程序退出
	<-forever
}
