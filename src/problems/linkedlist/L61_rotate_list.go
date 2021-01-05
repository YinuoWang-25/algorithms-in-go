package linkedlist

import "algorithms-in-go/src/problems/base"

func getLength(head *ListNode) int {
	length := 0
	for head != nil {
		length = length + 1
		head = (*ListNode)(head.Next)
	}
	return length
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	length := getLength(head)
	k = k % length

	dummy := &ListNode{Val: 0, Next: (*base.ListNode)(head)}
	slow := dummy
	fast := dummy
	for i := 0; i < k; i++ {
		fast = (*ListNode)(fast.Next)
	}
	for fast.Next != nil {
		fast = (*ListNode)(fast.Next)
		slow = (*ListNode)(slow.Next)
	}
	fast.Next = dummy.Next
	dummy.Next = slow.Next
	slow.Next = nil

	return (*ListNode)(dummy.Next)
}
