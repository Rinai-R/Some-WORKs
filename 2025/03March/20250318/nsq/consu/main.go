package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var consumer *nsq.Consumer
var err error
var logger, _ = zap.NewDevelopment()

func InitNSQ() {
	consumer, err = nsq.NewConsumer("test", "ch", nsq.NewConfig())
	if err != nil {
		logger.Panic("Failed to create consumer", zap.Error(err))
	}
}
func initLogger() {
	file, _ := os.OpenFile("./pkg/Logger/log/logger.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	writerSyncer := zapcore.AddSync(file)
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	logger = zap.New(core)
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
	initLogger()
	InitNSQ()
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		fmt.Println(string(message.Body))
		return nil
	}))
	if err := consumer.ConnectToNSQLookupd("localhost:4161"); err != nil {
		logger.Panic("Failed to connect to nsqlookupd", zap.Error(err))
	}
	select {}
}
