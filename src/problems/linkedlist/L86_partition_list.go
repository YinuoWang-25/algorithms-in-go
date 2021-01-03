package linkedlist

import "algorithms-in-go/src/problems/base"

func partition(head *ListNode, x int) *ListNode {
	less, more := &ListNode{Val: 0}, &ListNode{Val: 0}
	small, large := less, more
	for head != nil {
		if head.Val < x {
			small.Next = (*base.ListNode)(head)
			small = (*ListNode)(small.Next)
		} else {
			large.Next = (*base.ListNode)(head)
			large = (*ListNode)(large.Next)
		}
		head = (*ListNode)(head.Next)
	}
	large.Next = nil
	small.Next = more.Next
	return (*ListNode)(less.Next)
}
