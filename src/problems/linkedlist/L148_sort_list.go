package linkedlist

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	first := head
	// find middle node
	middle := findMiddleNode(head)
	second := middle.Next
	// split the list
	middle.Next = nil
	// sort both
	sortedFirst := sortList(first)
	sortedSecond := sortList((*ListNode)(second))

	// merge two list
	return mergeTwoLists(sortedFirst, sortedSecond)
}

func findMiddleNode(head *ListNode) *ListNode {
	// 1 -> 2 -> 3
	// 1 -> 2 -> 3 -> 4
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next

	for fast.Next != nil && fast.Next.Next != nil {
		slow = (*ListNode)(slow.Next)
		fast = fast.Next.Next
	}
	return slow
}
