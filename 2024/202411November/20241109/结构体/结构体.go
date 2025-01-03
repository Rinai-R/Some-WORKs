package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Number int
	Score  float64
}

func main() {
	var Class1 []Student
	var XiaoMing Student
	XiaoMing = Student{
		Name:   "XiaoMing",
		Age:    18,
		Number: 2024211885,
		Score:  100.0,
	}
	var LiHua Student
	LiHua = Student{
		Name:   "Lihua",
		Age:    18,
		Number: 2024211886,
		Score:  99.9,
	}
	Class1 = append(Class1, XiaoMing)
	Class1 = append(Class1, LiHua)

	fmt.Print(Class1)
	return
}
