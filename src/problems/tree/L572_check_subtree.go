package tree

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if t == nil {
		return true
	}
	if s.Val == t.Val && isSameTree(s, t) {
		return true
	}
	return isSubtree((*TreeNode)(s.Left), t) || isSubtree((*TreeNode)(s.Right), t)
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil || s.Val != t.Val {
		return true
	}
	return isSameTree((*TreeNode)(s.Left), (*TreeNode)(t.Left)) && isSameTree((*TreeNode)(s.Right), (*TreeNode)(t.Right))
}
