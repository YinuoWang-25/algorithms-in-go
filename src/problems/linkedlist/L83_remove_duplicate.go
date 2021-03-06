package linkedlist

import "algorithms-in-go/src/problems/base"

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: (*base.ListNode)(head)}
	prev := dummy

	for head != nil {
		last := head
		for last != nil && last.Val == head.Val {
			last = (*ListNode)(last.Next)
		}
		prev.Next = (*base.ListNode)(head)
		head.Next = (*base.ListNode)(last)
		prev = head
		head = (*ListNode)(prev.Next)
	}
	return (*ListNode)(dummy.Next)
}
