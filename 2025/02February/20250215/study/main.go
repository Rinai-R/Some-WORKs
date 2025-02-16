package main

import "container/heap"

func findKthLargest(nums []int, k int) int {
	h := &Heap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
	}
	for h.Len() > k {
		heap.Pop(h)
	}
	return heap.Pop(h).(int)
}

type Heap []int

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h Heap) Pop() any {
	Num := h[len(h)-1]
	h = h[:len(h)-1]
	return Num
}
