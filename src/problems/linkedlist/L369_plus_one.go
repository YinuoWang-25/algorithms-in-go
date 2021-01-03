package linkedlist

import "algorithms-in-go/src/problems/base"

func plusOne(head *ListNode) *ListNode {
	if head == nil {
		return &ListNode{Val: 1}
	}

	dummy := &ListNode{Val: 0}
	dummy.Next = (*base.ListNode)(head)

	slow, fast := dummy, dummy

	for fast != nil {
		if fast.Val != 9 {
			slow = fast
		}
		fast = (*ListNode)(fast.Next)
	}

	slow.Val = slow.Val + 1

	slow = (*ListNode)(slow.Next)
	for slow != nil {
		slow.Val = 0
		slow = (*ListNode)(slow.Next)
	}

	if dummy.Val == 1 {
		return dummy
	} else {
		return (*ListNode)(dummy.Next)
	}
}
