package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 10)

func Producer() {
	for i := 0; i < 10; i++ {
		var x int
		fmt.Scanln(&x)
		ch <- x
	}
	close(ch)
}

func Consumer() {
	for i := range ch {
		i--
		fmt.Println(i)
	}
}

func main() {
	go Producer()
	go Consumer()
	time.Sleep(time.Second * 10)
}
