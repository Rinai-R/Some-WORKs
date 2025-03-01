package main

import (
	"fmt"
	"sync"
)

type GoroutinePool struct {
	MaxGoroutineNum int

	Count sync.WaitGroup

	Tasks chan func()

	closed bool

	mutex sync.Mutex
}

func NewGoroutinePool(maxGoroutineNum int) *GoroutinePool {
	return &GoroutinePool{
		MaxGoroutineNum: maxGoroutineNum,
		Count:           sync.WaitGroup{},
		Tasks:           make(chan func(), maxGoroutineNum),
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
	pool.mutex.Lock()
	defer pool.mutex.Unlock()
	if pool.closed {
		fmt.Println("goroutine pool closed")
		return
	}
	pool.Count.Add(1)
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

func main() {
	pool := NewGoroutinePool(100)
	pool.Start()

	for i := 1; i <= 1000; i++ {
		x := i
		go pool.SubmitTask(func() {
			fmt.Println("这是第", x, "次执行任务！")
		})
	}
	defer pool.Wait()
	defer pool.Stop()
}
