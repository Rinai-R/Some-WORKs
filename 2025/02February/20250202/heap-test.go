package main

import (
	"container/heap"
	"fmt"
)

// IHeap 是一个最小堆的实现
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 方法将元素添加到堆中
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

// Pop 方法移除并返回堆顶元素
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// topKFrequent 函数找到数组中出现频率最高的 k 个元素
func topKFrequent(nums []int, k int) []int {
	// 统计每个数字的出现频率
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}

	// 创建最小堆
	h := &IHeap{}
	heap.Init(h)

	// 将元素推入堆并维护堆的大小
	for key, value := range m {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 从堆中提取结果
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

func main() {
	_ = make([]int, 0)
	nums := []int{1, 1, 216, 216, 216, 216, 216, 216, 6, 1, 2, 2, 3, 9, 9, 5, 6, 0, 6, 6, 9, 4, 5, 12, 6, 459, 15, 15, 216, 26, 15, 115, 15}
	k := 5
	fmt.Println(topKFrequent(nums, k))
}
