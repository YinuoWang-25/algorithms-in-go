package problems

import "algorithms-in-go/src/stack"

func PostorderTraversalI(root *TreeNode) []int {
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

		if cur.Left != nil {
			linkedStack.Push(cur.Left)
		}

		if cur.Right != nil {
			linkedStack.Push(cur.Right)
		}

	}
	reverse(res)
	return res
}

func reverse(res []int) {
	i := 0
	j := len(res) - 1

	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
}

/**
Recursive Solution
*/
func PostorderTraversalII(root *TreeNode) []int {
	var res []int
	postorderTraversalHelper(root, &res)
	return res
}

func postorderTraversalHelper(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postorderTraversalHelper((*TreeNode)(root.Left), res)
	postorderTraversalHelper((*TreeNode)(root.Right), res)
	*res = append(*res, root.Val)
}
