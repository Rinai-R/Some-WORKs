package main

import "fmt"

func main() {
	//s := make([]int, 0)
	//m := 1
	//s = append(s, m)
	//m = 2
	//fmt.Println(s)
	//l := make([][]int, 0)
	//l = append(l, s)
	//s[0] = 15
	//fmt.Println(l)

	s := make([]int, 0)
	for i := 0; i < 1024; i++ {
		s = append(s, i)
		fmt.Printf("%v\n", cap(s))
	}
}
