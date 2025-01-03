package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func food1(ch1 chan string) {
	for {
		time.Sleep(time.Second)
		ch1 <- "吃了超级大面包"
	}
}
func food2(ch2 chan string) {
	for {
		time.Sleep(time.Millisecond * 200)
		ch2 <- "吃了饺子"
	}
}
func food3(ch3 chan string) {
	for {
		time.Sleep(time.Millisecond * 600)
		ch3 <- "吃了牛子"
	}
}

func Eat(ch3 chan string, ch2 chan string, ch1 chan string) {
	for {
		select {
		case f := <-ch1:
			fmt.Println(f)
		case f := <-ch2:
			fmt.Println(f)
		case f := <-ch3:
			fmt.Println(f)
		}
		wg.Done()
	}
}

func main() {
	wg.Add(20)
	ch1 := make(chan string, 10)
	ch2 := make(chan string, 10)
	ch3 := make(chan string, 10)

	//食物生产线123
	go food1(ch1)
	go food2(ch2)
	go food3(ch3)

	go Eat(ch3, ch2, ch1)

	wg.Wait()

}
