package problems

import (
	"algorithms-in-go/src/problems/base"
	"algorithms-in-go/src/stack"
)

type TreeNode base.TreeNode

/**
Iterative Solution
*/

func PreorderTraversalI(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var linkedStack stack.LinkedStack

	linkedStack.Push(root)
	for {
		if linkedStack.IsEmpty() {
			break
		}
		cur := linkedStack.Pop().(*TreeNode)
		res = append(res, cur.Val)

		if cur.Right != nil {
			linkedStack.Push(cur.Right)
		}

		if cur.Left != nil {
			linkedStack.Push(cur.Left)
		}

	}
	return res
}

/**
Recursive Solution
*/
func PreorderTraversalII(root *TreeNode) []int {
	var res []int
	preorderTraversalHelper(root, &res)
	return res
}

func preorderTraversalHelper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorderTraversalHelper((*TreeNode)(root.Left), res)
	preorderTraversalHelper((*TreeNode)(root.Right), res)
}
