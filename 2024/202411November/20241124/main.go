package main

import "fmt"

type A interface {
	say()
}
type Ame struct {
}

type Chi struct {
}

func (s Ame) say() {
	fmt.Println("Hello")
}

func (s Chi) say() {
	fmt.Println("你好")
}

func (s Ame) Disco() {
	fmt.Println("Disco时间！")
}

func (s Chi) Niu() {
	fmt.Println("扭一下")
}

func exert(s A) {
	s.say()

	switch s.(type) {
	case Chi:
		ch := s.(Chi)
		ch.Niu()
	case Ame:
		am := s.(Ame)
		am.Disco()
	}
}

func main() {

	var C Chi
	var A Ame
	exert(C)
	exert(A)
	//st1 := Student1.Student{Name: "Rinai"}
	//st1.SetAge(155)
	//fmt.Println(st1.GetAge())
	//fmt.Println(st1)
	//
	//var x float64 = 5.6
	//var a A = x
	//fmt.Println(a)
	//var x1 string = "123"
	//var a1 A = x1
	//fmt.Println(a1)
	//var x2 string = "12443"
	//var a2 interface{} = x2
	//fmt.Println(a2)
	//return
}
