package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func write(ch1 chan int) {

	for i := 1; i <= 50; i++ {
		ch1 <- i*3 - 2
		time.Sleep(time.Millisecond * 50)
	}
	close(ch1)
}

func read(ch1 chan int) {
	defer wg.Done()
	for i := 1; i <= 50; i++ {
		h := <-ch1
		fmt.Println(h)
	}
}

func main() {
	wg.Add(50)
	ch1 := make(chan int, 60)
	go write(ch1)
	go read(ch1)

	wg.Wait()
}
