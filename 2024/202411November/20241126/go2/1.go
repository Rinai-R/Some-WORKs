package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		go func() {
			//fmt.Printf("第%v次输出", i)
			fmt.Println("hello world")
		}()
	}
	time.Sleep(time.Second * 2)

}
