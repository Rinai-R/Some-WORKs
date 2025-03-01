package test3

import (
	"sync"
	"testing"
	"time"
)

// 13.101
func TestOriginGoroutine(t *testing.T) {
	for i := 0; i < 50; i++ {
		cur := time.Now()
		wg := sync.WaitGroup{}
		wg.Add(1000000)
		for j := 0; j < 1000000; j++ {
			go func() {
				time.Sleep(time.Microsecond)
				wg.Done()
			}()
		}
		wg.Wait()
		t.Log(i, " ", time.Since(cur))
	}
}
