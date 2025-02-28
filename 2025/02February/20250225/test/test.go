package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j := len(nums)-1, len(nums)-2
	for j > 0 && nums[i] < nums[j] {
		i--
		j--
	}

	fmt.Println(j, i)
	firstBig := len(nums) - 1
	for nums[j] > nums[firstBig] {
		firstBig--
	}

	fmt.Println(nums[firstBig])
	nums[firstBig], nums[j] = nums[j], nums[firstBig]
	for l, r := i, len(nums)-1; l < r; {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func main() {
	cpuFile, _ := os.Create("./2025/02February/20250225/test/cpu.prof")
	defer cpuFile.Close()
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	// 生成内存分析文件（可选）
	memFile, _ := os.Create("。/2025/02February/20250225/test/mem.prof")
	defer memFile.Close()
	defer pprof.WriteHeapProfile(memFile)
	nextPermutation([]int{1, 5, 4, 3, 1, 4, 8, 5, 1, 3, 4, 7, 5, 1, 2, 34, 5, 8, 4, 3, 5, 4, 76, 3, 4, 12, 5, 48})
}
