package tree

func isValidBST(root *TreeNode) bool {
	return check(root, nil, nil)
}

func check(root *TreeNode, left *TreeNode, right *TreeNode) bool {
	if root == nil {
		return true
	}
	if left != nil && root.Val <= left.Val {
		return false
	}
	if right != nil && root.Val >= right.Val {
		return false
	}
	return check((*TreeNode)(root.Left), left, root) && check((*TreeNode)(root.Right), root, right)
}
