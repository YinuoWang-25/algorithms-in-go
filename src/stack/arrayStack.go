package stack

import (
	"sync"
)

type ArrayStack struct {
	arr  []interface{}
	size int
	lock sync.RWMutex
}

func (stack *ArrayStack) Push(v interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.arr = append(stack.arr, v)

	stack.size = stack.size + 1
}

func (stack *ArrayStack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("empty stack")
	}

	v := stack.arr[stack.size-1]

	newArray := make([]interface{}, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.arr[i]
	}
	stack.arr = newArray

	stack.size = stack.size - 1

	return v
}

func (stack *ArrayStack) Peek() interface{} {
	if stack.size == 0 {
		panic("empty stack")
	}

	v := stack.arr[stack.size-1]
	return v
}

func (stack *ArrayStack) Size() int {
	return stack.size
}

func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}
