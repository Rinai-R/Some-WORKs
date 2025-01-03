package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Score  []float32
	Number int
}

type Class struct {
	ClassRoom int
	ClassName string
	Student   []Student
}

func AddStudent(c *Class, s *Student) {
	c.Student = append(c.Student, *s)
}
func UpdateScore(s *Student, score float32) {
	s.Score = append(s.Score, score)
}

func CalculateAverage(s *Student) float32 {
	var Sum float32
	var i int
	i = 0
	for _, point := range s.Score {
		Sum += point
		i++
	}
	return Sum / float32(i)
}

func main() {
	var Class1 Class
	Class1 = Class{
		ClassRoom: 2011,
		ClassName: "Sunny Class",
		Student:   []Student{},
	}
	var XiaoMing Student
	XiaoMing = Student{
		Name:   "XiaoMing",
		Age:    18,
		Number: 2024211885,
	}
	var LiHua Student
	LiHua = Student{
		Name:   "Lihua",
		Age:    18,
		Number: 2024211886,
	}
	UpdateScore(&XiaoMing, 99.9)
	UpdateScore(&XiaoMing, 95.6)
	UpdateScore(&LiHua, 100)
	UpdateScore(&LiHua, 88.9)
	AddStudent(&Class1, &XiaoMing)
	AddStudent(&Class1, &LiHua)
	for _, Student := range Class1.Student {
		fmt.Println(Student.Name, Student.Age, Student.Number, CalculateAverage(&Student))
	}
}
