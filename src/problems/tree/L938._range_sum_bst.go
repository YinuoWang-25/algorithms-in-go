package tree

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	left := rangeSumBST((*TreeNode)(root.Left), low, high)
	right := rangeSumBST((*TreeNode)(root.Right), low, high)
	res := left + right
	if root.Val >= low && root.Val <= high {
		res += root.Val
	}
	return res
}
