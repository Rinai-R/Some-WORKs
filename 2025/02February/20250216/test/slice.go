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
}
