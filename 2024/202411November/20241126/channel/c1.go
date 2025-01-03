package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 4)
	ch1 <- 1
	ch1 <- 25
	close(ch1)
	for idx := 1; idx <= 2; idx++ {
		ac := <-ch1
		fmt.Println(ac)
	}
	return

}
