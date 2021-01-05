package linkedlist

import "algorithms-in-go/src/problems/base"

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = (*base.ListNode)(prev)
		prev = cur
		cur = (*ListNode)(next)
	}
	return prev
}
