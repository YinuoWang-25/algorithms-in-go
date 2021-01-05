package linkedlist

import "algorithms-in-go/src/problems/base"

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: (*base.ListNode)(head)}
	prev := dummy
	for i := 0; i < m-1; i++ {
		prev = (*ListNode)(prev.Next)
	}
	last := prev.Next
	cur := prev.Next
	for i := 0; i <= n-m; i++ {
		next := cur.Next
		cur.Next = prev.Next
		prev.Next = cur
		cur = next
	}
	last.Next = cur
	return (*ListNode)(dummy.Next)
}
