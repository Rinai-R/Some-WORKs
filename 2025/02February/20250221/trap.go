package main

import "fmt"

func trap(height []int) int {
	n := len(height)
	PreMax := make([]int, n+1)
	NeMax := make([]int, n+2)
	for i := 0; i < n; i++ {
		PreMax[i+1] = max(PreMax[i], height[i])
	}
	fmt.Println("PreMax", PreMax)
	for i := n - 1; i >= 0; i-- {
		NeMax[i+1] = max(NeMax[i+2], height[i])
	}
	fmt.Println("NeMax", NeMax)
	ans := 0
	for i := 0; i < n; i++ {
		ans += min(PreMax[i+1], NeMax[i+1]) - height[i]
	}
	return ans
}
func main() {
	trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
}
