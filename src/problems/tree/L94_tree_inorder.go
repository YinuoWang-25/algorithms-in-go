package tree

import (
	"algorithms-in-go/src/stack"
)

/**
Iterative Solution
*/
func InorderTraversalI(root *TreeNode) []int {
	var res []int

	var linkedStack stack.LinkedStack

	for {
		if root == nil {
			break
		}
		linkedStack.Push(root)
		root = (*TreeNode)(root.Left)
	}

	for {
		if linkedStack.IsEmpty() {
			break
		}
		cur := linkedStack.Pop().(*TreeNode)
		res = append(res, cur.Val)
		if cur.Right != nil {
			cur = (*TreeNode)(cur.Right)
			for {
				if cur == nil {
					break
				}
				linkedStack.Push(cur)
				cur = (*TreeNode)(cur.Left)
			}
		}
	}
	return res
}

/**
Recursive Solution
*/
func InorderTraversalII(root *TreeNode) []int {
	var res []int
	inorderTraversalHelper(root, &res)
	return res
}

func inorderTraversalHelper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorderTraversalHelper((*TreeNode)(root.Left), res)
	*res = append(*res, root.Val)
	inorderTraversalHelper((*TreeNode)(root.Right), res)
}
