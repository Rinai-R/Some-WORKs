package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	go func() {
		defer fmt.Println("goroutine")
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		panic("goroutine panic")

	}()
	time.Sleep(1 * time.Second)
	var b int
	//除以零的panic
	fmt.Scan(&b)
	a := 10 / b
	fmt.Println(a)
	//自定义panic
	panic("main panic")
	//不可到达的代码
	//recover令其他goroutine可以正常运行
	fmt.Println(time.Now())
}
