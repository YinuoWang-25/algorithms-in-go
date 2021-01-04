package linkedlist

import (
	"algorithms-in-go/src/priorityqueue"
	"algorithms-in-go/src/problems/base"
	"container/heap"
)

func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(priorityqueue.PriorityQueue, 0)

	for _, node := range lists {
		if node != nil {
			pq = append(pq, (*priorityqueue.ListNode)(node))
		}
	}

	if len(pq) == 0 {
		return nil
	}

	heap.Init(&pq)

	head := &ListNode{}
	dummyHead := head

	for len(pq) > 0 {
		min := heap.Pop(&pq)
		minNode := min.(*ListNode)
		head.Next = (*base.ListNode)(minNode)
		head = (*ListNode)(head.Next)

		if minNode.Next != nil {
			heap.Push(&pq, minNode.Next)
		}
	}
	return (*ListNode)(dummyHead.Next)
}
