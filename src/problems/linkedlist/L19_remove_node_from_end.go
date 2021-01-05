package linkedlist

import "algorithms-in-go/src/problems/base"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: (*base.ListNode)(head)}
	fast := dummy
	for i := 0; i < n; i++ {
		fast = (*ListNode)(fast.Next)
		if fast == nil {
			break
		}
	}
	slow := dummy
	for fast.Next != nil {
		slow = (*ListNode)(slow.Next)
		fast = (*ListNode)(fast.Next)
	}
	slow.Next = slow.Next.Next
	return (*ListNode)(dummy.Next)
}
