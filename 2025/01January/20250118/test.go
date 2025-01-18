package main

import (
	"fmt"
)

type user struct {
	name string
	age  uint64
}

func main() {
	u := []user{
		{"asong", 23},
		{"song", 19},
		{"asong2020", 18},
	}
	n := make([]*user, 0, len(u))
	for _, v := range u {
		n = append(n, &v)
	}
	fmt.Println(n)
	for _, v := range n {
		fmt.Println(v)
	}
}
