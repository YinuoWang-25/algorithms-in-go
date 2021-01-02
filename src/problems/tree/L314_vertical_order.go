package tree

import "math"

func verticalOrder(root *TreeNode) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	var queue []*TreeNode
	var colQueue []int
	// attention: need to use make instead of :=
	hash := make(map[int][]int)

	queue = append(queue, root)
	colQueue = append(colQueue, 0)

	min, max := 0, 0

	for len(queue) != 0 {
		curTreeNode := queue[0]
		curCol := colQueue[0]
		queue = queue[1:]
		colQueue = colQueue[1:]

		min = int(math.Min(float64(min), float64(curCol)))
		max = int(math.Max(float64(max), float64(curCol)))
		if _, ok := hash[curCol]; !ok {
			hash[curCol] = []int{}
		}
		hash[curCol] = append(hash[curCol], curTreeNode.Val)

		if curTreeNode.Left != nil {
			queue = append(queue, (*TreeNode)(curTreeNode.Left))
			colQueue = append(colQueue, curCol-1)
		}
		if curTreeNode.Right != nil {
			queue = append(queue, (*TreeNode)(curTreeNode.Right))
			colQueue = append(colQueue, curCol+1)
		}
	}
	for i := min; i <= max; i++ {
		res = append(res, hash[i])
	}
	return res
}
