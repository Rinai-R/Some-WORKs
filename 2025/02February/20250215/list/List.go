package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return Reverse(nil, head)
}

func Reverse(Prev *ListNode, Cur *ListNode) *ListNode {
	if Cur == nil {
		return Prev
	}
	Ne := Cur.Next
	Cur.Next = Prev
	return Reverse(Cur, Ne)
}
