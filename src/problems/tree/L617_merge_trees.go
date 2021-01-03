package tree

import "algorithms-in-go/src/problems/base"

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val = t1.Val + t2.Val

	t1.Left = (*base.TreeNode)(mergeTrees((*TreeNode)(t1.Left), (*TreeNode)(t2.Left)))
	t1.Right = (*base.TreeNode)(mergeTrees((*TreeNode)(t1.Right), (*TreeNode)(t2.Right)))
	return t1
}
