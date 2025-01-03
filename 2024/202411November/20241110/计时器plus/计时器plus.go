package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)
var StartTime time.Time
var ContinueTime time.Duration
var Run bool
var Flag []time.Duration

func Timer(ch chan int) {
	for {
		switch <-ch {
		case 1:
			if !Run {
				StartTime = time.Now()
				ContinueTime = 0
				Run = true
				fmt.Println("现在开始计时")
			} else {
				fmt.Println("计时已经开始了，别在摁1了")
			}
		case 0:
			if Run {
				ContinueTime += time.Now().Sub(StartTime)
				StartTime = time.Now()
				fmt.Println("计时已结束，时间为", ContinueTime)
				Run = false
			} else {
				fmt.Println("计时还没开始")
			}
		case 2:
			if Run {
				ContinueTime += time.Now().Sub(StartTime)
				Run = false
				fmt.Println("计时暂停了，已经运行了", ContinueTime)
			} else {
				StartTime = time.Now()
				Run = true
				fmt.Println("计时继续")
			}
		case 3: //3这个事件更加契合跑步计时的需求。
			if Run {
				ContinueTime += time.Now().Sub(StartTime)
				StartTime = time.Now()
				Flag = append(Flag, ContinueTime)
			} else {
				fmt.Println("还没开始呢")
			}
		}
	}
}

func cmd(ch chan int) {
	for {
		var cmd int
		fmt.Scan(&cmd)
		ch <- cmd
	}
}

func main() {

	go cmd(ch)

	go Timer(ch)
	time.Sleep(time.Second * 60)
	for i := 0; i < len(Flag); i++ {
		fmt.Printf("第%d个人所花费的时间为%v\n", i+1, Flag[i])
	}
}
