package main

import (
	"fmt"
)

func main() {
	_ = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9}
	x := []int{0, 0, 1, 2, 2, 2, 3}
	_ = []int{1, 1, 1, 1, 1, 1, 3, 3, 3, 3, 3, 8, 8, 9, 10, 10, 11, 14}
	_ = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	l := 0
	r := len(x) - 1
	num := -1
	for l < r {
		mid := (l + r) / 2
		if x[mid] < num {
			l = mid + 1
		} else {
			r = mid
		}
	}
	fmt.Println(x[l], l)

}
