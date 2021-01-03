package linkedlist

import "algorithms-in-go/src/problems/base"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	prev := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = (*base.ListNode)(l1)
			l1 = (*ListNode)(l1.Next)
		} else {
			prev.Next = (*base.ListNode)(l2)
			l2 = (*ListNode)(l2.Next)
		}
		prev = (*ListNode)(prev.Next)
	}
	if l1 != nil {
		prev.Next = (*base.ListNode)(l1)
	}
	if l2 != nil {
		prev.Next = (*base.ListNode)(l2)
	}
	return (*ListNode)(dummy.Next)
}
