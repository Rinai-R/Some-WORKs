package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func SortOddAndEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, oddTmp := head, head
	even, evenTmp := head.Next, head.Next
	for evenTmp != nil && evenTmp.Next != nil {
		oddTmp.Next = evenTmp.Next
		oddTmp = oddTmp.Next

		evenTmp.Next = oddTmp.Next
		evenTmp = evenTmp.Next
	}

	oddTmp.Next = nil

	even = ReverseList(even)
	return CombineListNode(odd, even)
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head
	for cur != nil {
		Ne := cur.Next
		cur.Next = prev
		prev = cur
		cur = Ne
	}
	return prev
}

func CombineListNode(list1, list2 *ListNode) *ListNode {
	tmp1 := list1
	tmp2 := list2
	dummy := &ListNode{}
	tmp := dummy
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val < tmp2.Val {
			tmp.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			tmp.Next = tmp2
			tmp2 = tmp2.Next
		}
		tmp = tmp.Next
	}
	if tmp1 != nil {
		tmp.Next = tmp1
	}
	if tmp2 != nil {
		tmp.Next = tmp2
	}
	return dummy.Next
}
