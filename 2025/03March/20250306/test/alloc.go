package main

import (
	"fmt"
	"runtime"
)

func main() {
	var x []int
	for i := 1; i <= 1000; i++ {
		x = make([]int, 0x3f3f3f3f)
	}
	fmt.Println(x)
	runtime.go
}
