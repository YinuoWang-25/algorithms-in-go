package linkedlist

import "algorithms-in-go/src/problems/base"

func oddEvenList(head *ListNode) *ListNode {
	odd, even := &ListNode{Val: 0}, &ListNode{Val: 0}
	first, second := odd, even
	place := 1
	for head != nil {
		if place%2 == 1 {
			first.Next = (*base.ListNode)(head)
			first = (*ListNode)(first.Next)
		} else {
			second.Next = (*base.ListNode)(head)
			second = (*ListNode)(second.Next)
		}
		head = (*ListNode)(head.Next)
		place = place + 1
	}
	second.Next = nil
	first.Next = even.Next
	return (*ListNode)(odd.Next)
}
