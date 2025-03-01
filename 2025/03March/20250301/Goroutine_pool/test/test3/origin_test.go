package test3

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 13.921
func TestOriginGoroutine(t *testing.T) {
	for i := 0; i < 50; i++ {
		cur := time.Now()
		wg := sync.WaitGroup{}
		//m := 0
		wg.Add(1000000)
		for j := 0; j < 1000000; j++ {
			go func() {
				fmt.Printf("")
				wg.Done()
			}()
		}
		wg.Wait()
		t.Log(i, " ", time.Since(cur))
	}
}
