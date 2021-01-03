package tree

/**
Top down
*/
func minDepthI(root *TreeNode) int {

	if root == nil {
		return 0
	}
	const MaxUint = ^uint(0)
	res := int(MaxUint >> 1)

	checkMinDepth(root, 1, &res)
	return res
}

func checkMinDepth(root *TreeNode, level int, global *int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if level < *global {
			*global = level
		}
	}
	checkMinDepth((*TreeNode)(root.Left), level+1, global)
	checkMinDepth((*TreeNode)(root.Right), level+1, global)
}

/**
Bottom Up
*/

func minDepthII(root *TreeNode) int {

	if root == nil {
		return 0
	}
	left, right := minDepthII((*TreeNode)(root.Left)), minDepthII((*TreeNode)(root.Right))
	if left == 0 || right == 0 {
		return left + right + 1
	}
	return min(left, right) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
