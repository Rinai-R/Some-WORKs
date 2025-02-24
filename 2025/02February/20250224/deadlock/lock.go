package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex
var lock2 sync.Mutex

func L2() {
	lock2.Lock()
	time.Sleep(1 * time.Second)
	lock.Lock()
	fmt.Println("lock2")
	lock.Unlock()
	lock2.Unlock()
}

func main() {
	lock.Lock()
	go L2()
	time.Sleep(1 * time.Second)

	lock2.Lock()
	fmt.Println("lock2")
	lock.Unlock()
	lock2.Unlock()
}
