package Test

import (
	"sync"
	"testing"
)

var mp1 = sync.Map{}
var wg1 = sync.WaitGroup{}

var mp = map[int]int{}
var lock = sync.RWMutex{}
var wg = sync.WaitGroup{}
var n = 20000

// func init() {
// 	for i := 0; i < n; i++ {
// 		mp1.Store(i, i)
// 	}
// 	for i := 0; i < n; i++ {
// 		mp[i] = i
// 	}
// }

func Test_syncmap(t *testing.T) {
	for i := 0; i < n; i++ {
		wg1.Add(1)
		go func() {
			for i := 0; i < n; i++ {
				mp1.Store(i, i)
			}
			wg1.Done()
		}()
	}
	wg1.Wait()
}

// func Test_lockmap(t *testing.T) {
// 	for i := 0; i < n; i++ {
// 		wg.Add(1)
// 		go func() {
// 			for i := 0; i < n; i++ {
// 				lock.RLock()
// 				_, _ = mp[i]
// 				lock.RUnlock()
// 			}
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }

// func Test(T *testing.T) {
// 	mp := sync.Map{}
// 	n := 100
// 	for i := 0; i < n; i++ {
// 		mp.Store(i, i)
// 	}
// 	for i := 0; i < n; i++ {
// 		fmt.Println(mp.Load(i))
// 	}
// }
