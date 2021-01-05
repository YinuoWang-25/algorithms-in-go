package linkedlist

import "algorithms-in-go/src/problems/base"

func reorderList(head *ListNode)  {
	if head == nil {
		return
	}
	first := head
	middle := findMiddle(head)
	second := middle.Next
	middle.Next = nil
	second = (*base.ListNode)(reverse((*ListNode)(second)))
	for first != nil && second != nil {
		nextFirst := first.Next
		nextSecond := second.Next
		first.Next = second
		second.Next = nextFirst
		first = (*ListNode)(nextFirst)
		second = nextSecond
	}
}
