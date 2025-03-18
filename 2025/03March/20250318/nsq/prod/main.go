package main

import (
	"encoding/json"
	"fmt"
	"runtime/debug"

	"github.com/nsqio/go-nsq"
)

var prod *nsq.Producer

func InitNSQ() {
	var err error
	config := nsq.NewConfig()

	prod, err = nsq.NewProducer("localhost:4150", config)
	if err != nil {
		panic(err)
	}
}

type Message struct {
	Msg string `json:"msg"`
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			// 使用 ANSI 代码将输出设置为红色（31 表示红色）
			fmt.Printf("\033[31mPanic: %v\n", r)
			// 输出堆栈信息，并在最后重置颜色（\033[0m）
			fmt.Printf("%s\033[0m", debug.Stack())
		}
	}()

	InitNSQ()
	topic := "test"
	msg := &Message{Msg: "Hello, World!"}
	bt, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	err = prod.Publish(topic, bt)
	if err != nil {
		panic(err)
	}
	info := fmt.Sprintf("Published message to %s", topic)
	fmt.Println(info)

}
