package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 1)
		close(ch)
	}() //如果去掉这个goroutine，会变成死锁，ch会阻塞main程序
	<-ch
	fmt.Println("main exit")
}
