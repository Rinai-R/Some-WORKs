package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Task struct {
	name     string
	duration time.Duration
}

var Cur time.Time

func (t *Task) run(done chan bool) {
	fmt.Printf("Task %s started in %v\n", t.name, time.Since(Cur))
	if x := <-done; x == true {
		fmt.Printf("Task %s finished in %v\n", t.name, time.Since(Cur))
	} else {
		return
	}
}

func roundRobinScheduler(tasks []Task, quantum time.Duration) {
	done := make(chan bool, len(tasks))

	for len(tasks) > 0 {
		// 处理当前任务
		task := tasks[0]
		tasks = tasks[1:]

		// 启动任务的 goroutine
		go task.run(done)
		if len(tasks) == 0 {
			time.Sleep(task.duration)
			done <- true
			return
		} else {
			// 模拟时间片轮转
			time.Sleep(quantum)
			task.duration -= quantum
			if task.duration > 0 {
				done <- false
				tasks = append(tasks, task)
			} else {
				done <- true
				continue
			}
		}
	}
}

func main() {
	Cur = time.Now()
	tasks := []Task{
		{"Task 1", 2 * time.Second},
		{"Task 2", 3 * time.Second},
		{"Task 3", 1 * time.Second},
	}

	quantum := 100 * time.Millisecond
	roundRobinScheduler(tasks, quantum)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-sigChan:
		// 捕获到中断信号
		fmt.Printf("Received signal: %s. Exiting...\n", sig)
		// 执行优雅的关闭操作（比如关闭任务、清理资源等）
		// 这里可以根据需要添加你的清理逻辑
	}
}
