package linkedlist

import "algorithms-in-go/src/problems/base"

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{Val: 0, Next: nil}
	cur, pre := head, newHead
	for cur != nil {
		next := cur.Next
		for pre.Next != nil && pre.Next.Val < cur.Val {
			pre = (*ListNode)(pre.Next)
		}
		cur.Next = pre.Next
		pre.Next = (*base.ListNode)(cur)
		pre = newHead
		cur = (*ListNode)(next)
	}
	return (*ListNode)(newHead.Next)
}
