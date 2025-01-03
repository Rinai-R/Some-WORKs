package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var ans int

func add(n int) {
	for i := 2; i <= n; i++ {
		ans *= i
	}
}

func main() {

	for i := 2; i <= 10; i++ {
		wg.Add(1)
		ans = 1
		go add(i)
		fmt.Println(ans)
		wg.Done()
	}
	wg.Wait()
}
