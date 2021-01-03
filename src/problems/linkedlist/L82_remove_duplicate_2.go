package linkedlist

import "algorithms-in-go/src/problems/base"

func deleteDuplicatesII(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: (*base.ListNode)(head)}
	prev := dummy

	for head != nil {
		last := head
		for last != nil && last.Val == head.Val {
			last = (*ListNode)(last.Next)
		}
		if last == (*ListNode)(head.Next) {
			prev = head
		} else {
			prev.Next = (*base.ListNode)(last)
		}
		head = last
	}
	return (*ListNode)(dummy.Next)
}
