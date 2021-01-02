package tree

import "algorithms-in-go/src/problems/base"

type Node base.NarayTreeNode

func naryLevelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue []*Node

	queue = append(queue, root)

	for len(queue) != 0 {
		var cur []int
		size := len(queue)
		for i := 0; i < size; i++ {
			curTreeNode := queue[0]
			queue = queue[1:]
			cur = append(cur, curTreeNode.Val)

			if curTreeNode.Children != nil {
				for i := 0; i < len(curTreeNode.Children); i++ {
					child := curTreeNode.Children[i]
					if child != nil {
						queue = append(queue, (*Node)(child))
					}
				}
			}
		}
		res = append(res, cur)
	}

	return res
}
