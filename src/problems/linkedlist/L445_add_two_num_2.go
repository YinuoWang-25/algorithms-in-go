package linkedlist

import "algorithms-in-go/src/problems/base"

func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	var tail *ListNode

	var s1, s2 []int
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = (*ListNode)(l1.Next)
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = (*ListNode)(l2.Next)
	}

	sum := 0
	for len(s1) > 0 || len(s2) > 0 || sum > 0 {
		if len(s1) > 0 {
			sum = sum + s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			sum = sum + s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		curNode := &ListNode{Val: sum % 10}
		sum = sum / 10
		if tail != nil {
			curNode.Next = (*base.ListNode)(tail)
		}
		tail = curNode
	}
	return tail
}
