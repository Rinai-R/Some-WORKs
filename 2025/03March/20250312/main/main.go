package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())

	_, err := c.AddFunc("*/1 * * * * *", func() {
		fmt.Println("任务执行时间:", time.Now().Format("2006-01-02 15:04:05"))
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Start()

	select {}
}
