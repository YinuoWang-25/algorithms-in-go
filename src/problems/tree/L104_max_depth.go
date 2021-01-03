package tree

import (
	"math"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(1 + math.Max(float64(maxDepth((*TreeNode)(root.Left))), float64(maxDepth((*TreeNode)(root.Right)))))
}
