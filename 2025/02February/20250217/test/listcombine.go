package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type ListHeap []*ListNode

func (l *ListHeap) Len() int {
	return len(*l)
}

func (l *ListHeap) Less(i, j int) bool {
	return (*l)[i].Val < (*l)[j].Val
}

func (l *ListHeap) Swap(i, j int) {
	(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
}

func (l *ListHeap) Push(x any) {
	*l = append(*l, x.(*ListNode))
}

func (l *ListHeap) Pop() any {
	res := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return res
}
