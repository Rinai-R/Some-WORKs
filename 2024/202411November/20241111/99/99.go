package main

import "fmt"

func main() {
	var a, b int
	a = 1
	for a <= 9 {
		b = 1
		for b <= a {
			fmt.Printf("%d * %d = %d  ", a, b, a*b)
			b++
		}
		fmt.Println()
		a++
	}
}
