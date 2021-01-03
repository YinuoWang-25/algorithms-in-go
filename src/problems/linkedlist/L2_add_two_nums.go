package linkedlist

import "algorithms-in-go/src/problems/base"

type ListNode base.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	sum, cur := 0, dummy
	for l1 != nil || l2 != nil || sum != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = (*ListNode)(l1.Next)
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = (*ListNode)(l2.Next)
		}
		sum = sum + n1 + n2
		cur.Next = (*base.ListNode)(&ListNode{Val: sum % 10})
		cur = (*ListNode)(cur.Next)
		sum = sum / 10
	}
	return (*ListNode)(dummy.Next)
}
