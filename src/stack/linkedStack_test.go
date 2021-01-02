package stack

import (
	"fmt"
	"testing"
)

func TestLinkedStack(t *testing.T) {
	arrayStack := new(ArrayStack)
	arrayStack.Push(1)
	arrayStack.Push(2)
	arrayStack.Push(3)
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Push(4)
	fmt.Println("pop:", arrayStack.Pop())
}