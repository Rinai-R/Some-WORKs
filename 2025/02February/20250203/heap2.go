package main

import "container/heap"

type MaxHeap []int
type MinHeap []int

type MedianFinder struct {
	*MaxHeap
	*MinHeap
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Less(i, j int) bool {
	return h[j] < h[i]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Constructor() MedianFinder {
	ah := &MaxHeap{}
	heap.Init(ah)
	ih := &MinHeap{}
	heap.Init(ih)
	return MedianFinder{
		MaxHeap: ah,
		MinHeap: ih,
	}
}

func (this *MedianFinder) AddNum(num int) {
	switch this.MaxHeap.Len() - this.MinHeap.Len() {
	case 0:
		if num > (*this.MaxHeap)[0] {
			heap.Push(this.MaxHeap, num)
		} else {
			heap.Push(this.MinHeap, num)
		}
	case -1:
		if num > (*this.MaxHeap)[0] {
			heap.Push(this.MaxHeap, num)
		} else {
			heap.Push(this.MaxHeap, heap.Pop(this.MinHeap))
			heap.Push(this.MinHeap, num)
		}

	}
}

func (this *MedianFinder) FindMedian() float64 {
	switch this.MaxHeap.Len() - this.MinHeap.Len() {
	case 0:
		return float64((*this.MaxHeap)[0]-(*this.MinHeap)[0]) / 2.0
	case 1:
		return float64((*this.MaxHeap)[0])
	default:
		return float64((*this.MinHeap)[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
