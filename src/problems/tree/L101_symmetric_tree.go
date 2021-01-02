package tree

/**
Recursive Solution
*/
func isSymmetricI(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSymmetric((*TreeNode)(root.Left), (*TreeNode)(root.Right))
}

func checkSymmetric(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return checkSymmetric((*TreeNode)(p.Left), (*TreeNode)(q.Right)) && checkSymmetric((*TreeNode)(p.Right), (*TreeNode)(q.Left))
}

/**
Iterative Solution
*/
func isSymmetricII(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue []*TreeNode
	queue = append(queue, (*TreeNode)(root.Left))
	queue = append(queue, (*TreeNode)(root.Right))
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
			queue = append(queue, (*TreeNode)(second.Right))
			queue = append(queue, (*TreeNode)(second.Left))
			queue = append(queue, (*TreeNode)(first.Right))
		}
	}
	return true
}
