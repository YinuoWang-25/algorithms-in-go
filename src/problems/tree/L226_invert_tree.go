package tree

import "algorithms-in-go/src/problems/base"

/**
Recursive Solution
*/
func invertTreeI(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTreeI((*TreeNode)(root.Right))
	right := invertTreeI((*TreeNode)(root.Left))
	root.Left = (*base.TreeNode)(left)
	root.Right = (*base.TreeNode)(right)
	return root
}

/**
Iterative Solution
*/
func invertTreeII(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		left := cur.Left
		right := cur.Right
		cur.Left = right
		cur.Right = left
		if left != nil {
			queue = append(queue, (*TreeNode)(left))
		}
		if right != nil {
			queue = append(queue, (*TreeNode)(right))
		}
	}
	return root
}
