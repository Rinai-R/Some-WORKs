package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type GoroutinePool struct {
	//最大并发量
	MaxGoroutineNum int
	//当前排队数量
	Count sync.WaitGroup
	//待完成的工作管道
	Tasks chan func()
	//管道的状态
	closed bool
	//锁
	mutex sync.Mutex
}

func NewGoroutinePool(maxGoroutineNum, MaxQueueNum int) *GoroutinePool {
	return &GoroutinePool{
		MaxGoroutineNum: maxGoroutineNum,
		Count:           sync.WaitGroup{},
		Tasks:           make(chan func(), MaxQueueNum),
		closed:          true,
		mutex:           sync.Mutex{},
	}
}

func (pool *GoroutinePool) Start() {
	pool.closed = false
	for i := 0; i < pool.MaxGoroutineNum; i++ {
		go pool.Work()
	}
}

func (pool *GoroutinePool) Work() {
	for task := range pool.Tasks {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("goroutine panic:", r)
				}
				pool.Count.Done()
			}()
			task()
		}()
	}
}

func (pool *GoroutinePool) SubmitTask(task func()) {
	//pool.mutex.Lock()
	//defer pool.mutex.Unlock()
	pool.Count.Add(1)
	//if pool.closed {
	//	fmt.Println("goroutine pool closed")
	//	return
	//}
	pool.Tasks <- task
}

func (pool *GoroutinePool) Stop() {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	pool.closed = true
	close(pool.Tasks)
}

func (pool *GoroutinePool) Wait() {
	pool.Count.Wait()
}

func Test_Pool(t *testing.T) {
	for j := 0; j < 50; j++ {
		pool := NewGoroutinePool(1000000, 1000000)
		pool.Start()
		cur := time.Now()
		for i := 0; i < 1000000; i++ {
			pool.SubmitTask(func() {
				time.Sleep(time.Microsecond)
			})
		}
		pool.Wait()
		pool.Stop()
		fmt.Printf("%d goroutines cost %v\n", j, time.Since(cur))
		//if m != 1000 {
		//	t.Error("Expected 1000 tasks but got ", m)
		//} else {
		//	fmt.Println("OK")
		//}
	}
}
