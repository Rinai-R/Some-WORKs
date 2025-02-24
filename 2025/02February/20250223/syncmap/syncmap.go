package main

import (
	"fmt"
	"sync"
)

type mystruct struct {
	name string
	age  int
}

func main() {
	var sm sync.Map

	// Store存储
	sm.Store("foo", 42)
	sm.Store("bar", "Hello, World!")

	// Load可以检查元素是否存在，也可以存入元素
	if value, ok := sm.Load("foo"); ok {
		fmt.Println("foo:", value)
	}

	sm.Delete("foo")

	if _, ok := sm.Load("foo"); !ok {
		fmt.Println("foo not found")
	}

	sm.Store(&mystruct{name: "foo", age: 42}, mystruct{name: "bar", age: 42})

	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
