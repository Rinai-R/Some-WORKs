package main

import (
	"fmt"
	"time"
)

type Task interface {
	Execute()
}

type PrintTask struct {
	thing string
}
type CalculationTask struct {
	A int
	B int
}
type SleepTask struct {
	Duration int
}

func (c PrintTask) Execute() {
	fmt.Println(c.thing)
}
func (c CalculationTask) Execute() {
	fmt.Println(c.A + c.B)
}
func (c SleepTask) Execute() {
	time.Sleep(time.Duration(c.Duration) * time.Second)
}

type Scheduler struct {
	Tasks []Task
}

func (s *Scheduler) AddTask(task Task) {
	s.Tasks = append(s.Tasks, task)
}

func (s *Scheduler) RunAll() {
	for _, t := range s.Tasks {
		t.Execute()
	}
}

func main() {
	greet := PrintTask{thing: "Good Morning!"}
	CountNum := CalculationTask{A: 22, B: 19}
	Sleep := SleepTask{Duration: 2}
	var s1 Scheduler
	s1.AddTask(greet)
	s1.AddTask(Sleep)
	s1.AddTask(CountNum)
	s1.RunAll()
	return
}
