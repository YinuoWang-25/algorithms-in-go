package linkedlist

import "algorithms-in-go/src/problems/base"

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Val: 0, Next: (*base.ListNode)(head)}
	prev := dummy
	for head != nil {
		if head.Val == val {
			prev.Next = head.Next
		}
		prev = (*ListNode)(prev.Next)
		head = (*ListNode)(prev.Next)
	}
	return (*ListNode)(dummy.Next)
}
