package main

import "fmt"

func main() {
	var x []int
	x = append(x, 1, 2, 3)
	xm(&x)
	fmt.Println(x)
}

func xm(x *[]int) {
	fmt.Println(&x)
	*x = append(*x, 1, 2, 3)
}
