package tree

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue []*TreeNode

	queue = append(queue, root)

	for len(queue) != 0 {
		var cur []int
		size := len(queue)
		for i := 0; i < size; i++ {
			curTreeNode := queue[0]
			queue = queue[1:]
			cur = append(cur, curTreeNode.Val)
			if curTreeNode.Left != nil {
				queue = append(queue, (*TreeNode)(curTreeNode.Left))
			}
			if curTreeNode.Right != nil {
				queue = append(queue, (*TreeNode)(curTreeNode.Right))
			}
		}
		res = append(res, cur)
	}
	return res
}
