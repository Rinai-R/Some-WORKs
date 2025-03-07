package main

import (
	"fmt"
	"sync"
)

var o sync.Once

func hello() {
	fmt.Println("hello")
}
func fun() {
	fmt.Println("fun")
}

func main() {
	o.Do(hello)
	o.Do(hello)
	o.Do(fun)
}
