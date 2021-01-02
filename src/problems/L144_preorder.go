package problems

type LinkedNode struct {
	Val  *TreeNode
	Next *LinkedNode
}

type LinkedStack struct {
	head *LinkedNode
	size int
}

func (stack *LinkedStack) Push(v *TreeNode) {
	newHead := &LinkedNode{v, stack.head}

	stack.head = newHead

	stack.size = stack.size + 1
}

func (stack *LinkedStack) Pop() *TreeNode {

	if stack.size == 0 {
		panic("empty stack")
	}

	res := stack.head.Val

	stack.head = stack.head.Next

	stack.size = stack.size - 1

	return res
}

func (stack *LinkedStack) IsEmpty() bool {
	return stack.size == 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func iterative(root *TreeNode) []int {

	var res []int
	if root == nil {
		return res
	}

	var stack LinkedStack
	stack.Push(root)
	for {
		if stack.IsEmpty(){
			break
		}
		cur := stack.Pop()
		res = append(res, cur.Val)

		if cur.Right != nil {
			stack.Push(cur.Right)
		}

		if cur.Left != nil {
			stack.Push(cur.Left)
		}

	}
	return res
}


func recursive(root *TreeNode) []int{
	var res []int
	recursiveHelper(root, &res)
	return res
}

func recursiveHelper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	recursiveHelper(root.Left,res)
	recursiveHelper(root.Right,res)
}