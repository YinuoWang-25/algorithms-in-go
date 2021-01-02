package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	deque := []*TreeNode{root}
	level := 0
	for len(deque) > 0 {
		size := len(deque)
		curLevel := make([]int, 0, 0)
		if level%2 == 0 {
			for i := size - 1; i >= 0; i-- {
				curNode := deque[len(deque)-1]
				curLevel = append(curLevel, curNode.Val)
				deque = deque[:len(deque)-1]
				if curNode.Left != nil {
					deque = append([]*TreeNode{(*TreeNode)(curNode.Left)}, deque...)
				}
				if curNode.Right != nil {
					deque = append([]*TreeNode{(*TreeNode)(curNode.Right)}, deque...)
				}
			}
		} else {
			for i := 0; i < size; i++ {
				curNode := deque[0]
				curLevel = append(curLevel, curNode.Val)
				deque = deque[1:]
				if curNode.Right != nil {
					deque = append(deque, (*TreeNode)(curNode.Right))
				}
				if curNode.Left != nil {
					deque = append(deque, (*TreeNode)(curNode.Left))
				}
			}
		}
		res = append(res, curLevel)
		level = level + 1
	}

	return res
}
