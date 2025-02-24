package main

import (
	"fmt"
	"sort"
)

func main() {
	x := [][]int{{1, 4}, {0, 4}}
	sort.SliceIsSorted([][]int{{1, 4}, {0, 4}}, func(i, j int) bool {
		return x[i][0] < x[j][0]
	})
	fmt.Println(x)

}
