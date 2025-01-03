package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var ans int = 1
var wg sync.WaitGroup

func fact(n int) {
	lock.Lock()
	ans = 1
	for i := n; i >= 1; i-- {
		ans *= i
	}
	lock.Unlock()
}

func main() {
	for i := 2; i <= 10; i++ {
		wg.Add(1)

		go fact(i)
		go func() {
			lock.Lock()
			fmt.Println(ans)
			lock.Unlock()
			wg.Done()
		}()

	}

	wg.Wait()

}
