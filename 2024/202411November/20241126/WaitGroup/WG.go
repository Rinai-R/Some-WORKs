package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		func(i int) {
			wg.Done()
			fmt.Printf("这是第%v次输出~\n", i)
		}(i)
	}
	wg.Wait()
}
