package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"time"
)

func main() {
	err := sentinel.InitDefault()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "Test",
			Threshold:              10,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	for i := 0; i <= 1000; i++ {
		go func() {

			// 使用资源 "test" 进行限流控制
			e, b := sentinel.Entry("Test")
			if b != nil {
				// 如果返回错误，表示该请求被限流
				fmt.Println("限流了！！！")
			} else {
				// 请求未被限流，处理业务逻辑
				fmt.Println("----")
				e.Exit()
			}
		}()
	}

	// 阻塞主线程，等待 goroutines 完成
	time.Sleep(3 * time.Second)
}
