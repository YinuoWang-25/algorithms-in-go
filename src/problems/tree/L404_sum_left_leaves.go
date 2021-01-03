package tree

func sumOfLeftLeaves(root *TreeNode) int {
	var res int

	if root == nil {
		return 0
	}

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			res += root.Left.Val
		} else {
			res = res + sumOfLeftLeaves((*TreeNode)(root.Left))
		}
	}

	if root.Right != nil {
		res = res + sumOfLeftLeaves((*TreeNode)(root.Right))
	}

	return res
}
