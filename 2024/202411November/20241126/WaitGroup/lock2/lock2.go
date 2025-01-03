package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var lock sync.RWMutex

func read() {
	defer wg.Done()
	lock.RLock()
	fmt.Println("读取中....")
	time.Sleep(time.Second)
	fmt.Println("读入成功！！！")
	lock.RUnlock()
}

func write() {
	lock.Lock()
	fmt.Println("写入中....")
	time.Sleep(time.Second * 4)
	fmt.Println("成功写入了~")
	lock.Unlock()
}
func main() {
	wg.Add(6)
	for i := 1; i <= 6; i++ {
		go read()
	}

	go write()

	wg.Wait()
}
