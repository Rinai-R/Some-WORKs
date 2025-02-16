package main

import "fmt"

func main() {
	s := make([]int, 0)
	s = append(s, 1, 2)
	fmt.Println(s[0])
	p := &s[0]
	s = s[:len(s)-2]
	s = make([]int, 2)
	fmt.Println(*p)

	m := X{}
	(&m).Set(1)
	fmt.Println(Test(&m))

}

type L interface {
	Get() int
	Set(int)
}

type X struct {
	x int
}

func (x *X) Get() int {
	return (*x).x
}

func (x *X) Set(i int) {
	(*x).x = i
}

func Test(l L) int {
	return l.Get()
}
