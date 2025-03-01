package main

import (
	"fmt"
	"sync"
)

type GoroutinePool struct {
	MaxGoroutineNum int

	Count sync.WaitGroup

	Tasks chan func()
}

func NewGoroutinePool(maxGoroutineNum int) *GoroutinePool {
	return &GoroutinePool{
		MaxGoroutineNum: maxGoroutineNum,
		Count:           sync.WaitGroup{},
		Tasks:           make(chan func(), maxGoroutineNum),
	}
}

func (pool *GoroutinePool) Start() {
	for i := 0; i < pool.MaxGoroutineNum; i++ {
		go pool.Work()
	}
}

func (pool *GoroutinePool) Work() {
	for task := range pool.Tasks {
		task()
		pool.Count.Done()
	}
}

func (pool *GoroutinePool) SubmitTask(task func()) {
	pool.Count.Add(1)
	pool.Tasks <- task
}

func (pool *GoroutinePool) Stop() {
	close(pool.Tasks)
}

func (pool *GoroutinePool) Wait() {
	pool.Count.Wait()
}

func main() {
	pool := NewGoroutinePool(500)
	pool.Start()
	for i := 1; i <= 1000; i++ {
		x := i
		pool.SubmitTask(func() {
			fmt.Println("这是第", x, "次执行任务！")
		})
	}
	defer pool.Stop()
	defer pool.Wait()
}
