package tree

/**
Recursive Solution
*/
func isSameTreeI(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTreeI((*TreeNode)(p.Left), (*TreeNode)(q.Left)) && isSameTreeI((*TreeNode)(p.Right), (*TreeNode)(q.Right))
}

/**
Iterative Solution
*/
func isSameTreeII(p *TreeNode, q *TreeNode) bool {
	var queue []*TreeNode
	queue = append(queue, p)
	queue = append(queue, q)
	for len(queue) > 0 {
		first := queue[0]
		second := queue[1]
		queue = queue[2:]
		if first == nil && second == nil {
			continue
		} else if first == nil || second == nil || first.Val != second.Val {
			return false
		} else {
			queue = append(queue, (*TreeNode)(first.Left))
			queue = append(queue, (*TreeNode)(second.Left))
			queue = append(queue, (*TreeNode)(first.Right))
			queue = append(queue, (*TreeNode)(second.Right))
		}
	}
	return true
}
