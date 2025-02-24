package main

import "fmt"

func main() {
	var x []int
	x = append(x, 1, 3, 5, 7, 5)
	m(x)
	fmt.Println(x)

}

func m(x []int) {
	x[1] = 5
	fmt.Println(5 / 2)
}
