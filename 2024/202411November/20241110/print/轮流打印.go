package main

import (
	"fmt"
	"time"
)

var flag int

func PrintOdd() {
	for flag <= 100 {
		if flag%2 == 1 {
			fmt.Println(flag, 0)
			flag++
		}
	}
}
func PrintEven() {
	for flag <= 100 {
		if flag%2 == 0 {
			fmt.Println(flag, 1)
			flag++
		}
	}
}

func main() {
	flag = 1
	go PrintOdd()
	go PrintEven()
	time.Sleep(time.Second * 5)
}
