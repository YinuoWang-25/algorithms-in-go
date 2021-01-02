package stack

import "sync"

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

type LinkedStack struct {
	head *LinkedNode
	size int
	lock sync.RWMutex
}

func (stack *LinkedStack) Push(v int) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	newHead := &LinkedNode{v, stack.head}

	stack.head = newHead

	stack.size = stack.size + 1
}

func (stack *LinkedStack) Pop() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("empty stack")
	}

	res := stack.head.Val

	stack.head = stack.head.Next

	return res
}

func (stack *LinkedStack) Peek() int {
	if stack.size == 0 {
		panic("empty stack")
	}

	return stack.head.Val
}

func (stack *LinkedStack) Size() int {
	return stack.size
}

func (stack *LinkedStack) IsEmpty() bool {
	return stack.size == 0
}
