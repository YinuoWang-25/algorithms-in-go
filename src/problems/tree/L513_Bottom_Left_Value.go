package tree

func findBottomLeftValue(root *TreeNode) int {
	var res int
	counted := 0
	findBottomLeftValueTraverse(root, &res, 1, &counted)
	return res
}

func findBottomLeftValueTraverse(root *TreeNode, res *int, level int, counted *int) {
	if root == nil {
		return
	}
	if level > *counted {
		*res = root.Val
		*counted = *counted + 1
	}
	findBottomLeftValueTraverse((*TreeNode)(root.Left), res, level + 1, counted)
	findBottomLeftValueTraverse((*TreeNode)(root.Right), res, level + 1, counted)
}