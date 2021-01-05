package linkedlist

import "algorithms-in-go/src/problems/base"

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	first := head
	middle := findMiddle(head)
	second := middle.Next
	middle.Next = nil
	second = (*base.ListNode)(reverse((*ListNode)(second)))
	for first != nil && second != nil {
		if first.Val != second.Val {
			return false
		}
		first = (*ListNode)(first.Next)
		second = second.Next
	}
	if first != nil {
		first = (*ListNode)(first.Next)
	}
	return first == nil && second == nil
}

func findMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = (*ListNode)(slow.Next)
		fast = (*ListNode)(fast.Next.Next)
	}
	return slow
}
