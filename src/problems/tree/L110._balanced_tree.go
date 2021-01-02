package tree

import "algorithms-in-go/src/utils"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := getHeight((*TreeNode)(root.Left))
	right := getHeight((*TreeNode)(root.Right))
	if utils.Abs(left-right) > 1 {
		return false
	}
	return isBalanced((*TreeNode)(root.Left)) && isBalanced((*TreeNode)(root.Right))
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getHeight((*TreeNode)(root.Left))
	right := getHeight((*TreeNode)(root.Right))
	var higher = utils.Max(left, right)
	return higher + 1
}
