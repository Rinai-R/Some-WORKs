package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	f := make([][]int, n)
	path := make([][][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, 2)
		path[i] = make([][]int, 2)
	}
	f[0][1] = nums[0]
	path[0][1] = append([]int{}, nums[0])

	for i := 1; i < n; i++ {
		if f[i-1][0] > f[i-1][1] {
			f[i][0] = f[i-1][0]
			path[i][0] = append([]int{}, path[i-1][0]...)
		} else {
			f[i][0] = f[i-1][1]
			path[i][0] = append([]int{}, path[i-1][1]...)
		}

		f[i][1] = f[i-1][0] + nums[i]
		cur := append([]int{}, path[i-1][0]...)
		cur = append(cur, nums[i])

		path[i][1] = cur
	}
	if f[n-1][0] > f[n-1][1] {
		fmt.Println(path[n-1][0])
		return f[n-1][0]
	} else {
		fmt.Println(path[n-1][1])
		return f[n-1][1]
	}
}

func main() {
	fmt.Println(rob([]int{1, 1, 1, 2}))
}
