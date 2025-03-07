package main

import (
	"fmt"
	"runtime"
)

func fun() *int {
	i := 1
	return &i
}

func main() {
	fmt.Println(*fun())
	runtime.Breakpoint()
}
