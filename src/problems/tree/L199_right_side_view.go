package tree

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0, 0)
	rightSideViewTraverse(root, &res, 1) // pass pointer to res in helper function
	return res
}

func rightSideViewTraverse(root *TreeNode, res *[]int, level int) {
	if root == nil {
		return
	}
	if level > len(*res) {
		*res = append(*res, root.Val)
	}
	rightSideViewTraverse((*TreeNode)(root.Right), res, level + 1)
	rightSideViewTraverse((*TreeNode)(root.Left), res, level + 1)
}


