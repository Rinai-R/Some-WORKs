package main

import "fmt"

func numSquares(n int) int {
	var f []int
	var nums []int
	for i := 1; i <= n/i; i++ {
		nums = append(nums, i*i)
	}

	Len := len(nums) - 1
	f = append(f, 0)
	for i := 1; i <= n; i++ {
		MinVal := 10000
		for j := Len; j >= 0; j-- {
			if nums[j] > i {
				continue
			}
			var k int
			for k = 1; k*nums[j] <= i; k++ {
			}
			k--
			MinVal = min(MinVal, f[i-k*nums[j]]+k)
		}
		f = append(f, MinVal)
	}
	fmt.Printf("\n")
	fmt.Println(f)
	return f[n]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func main() {
	numSquares(12)
}
