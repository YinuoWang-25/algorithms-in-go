package linkedlist

import "algorithms-in-go/src/problems/base"

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: 0}
	dummy.Next = (*base.ListNode)(head)
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		first, second := prev.Next, prev.Next.Next
		first.Next = second.Next
		second.Next = first
		prev.Next = second
		prev = (*ListNode)(first)
	}
	return (*ListNode)(dummy.Next)
}
