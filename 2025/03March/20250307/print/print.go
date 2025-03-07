package main

import (
	"fmt"
	"time"
)

var str = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y"}

func main() {
	var i int
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)
	defer close(ch1)
	defer close(ch2)
	ch1 <- struct{}{}
	go func() {
		for {
			<-ch1
			if i >= len(str) {
				break
			}
			fmt.Println(str[i], 1)
			i++
			ch2 <- struct{}{}

		}
	}()

	go func() {
		for {
			<-ch2
			if i >= len(str) {
				break
			}
			fmt.Println(str[i], 2)
			i++
			ch1 <- struct{}{}

		}
	}()
	time.Sleep(time.Second * 100)
}
