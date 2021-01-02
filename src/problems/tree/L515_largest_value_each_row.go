package tree

func largestValues(root *TreeNode) []int {
	var res []int
	largestValuesTraverse(root, &res, 0)
	return res
}

func largestValuesTraverse(root *TreeNode, res *[]int, level int) {
	if root == nil {
		return
	}
	if level >= len(*res) {
		*res = append(*res, root.Val)
	} else if root.Val > (*res)[level] {
		(*res)[level] = root.Val
	}
	largestValuesTraverse((*TreeNode)(root.Left), res, level+1)
	largestValuesTraverse((*TreeNode)(root.Right), res, level+1)
}
