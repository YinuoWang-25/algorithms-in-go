package tree

func maxDepthForNary(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Children == nil || len(root.Children) == 0 {
		return 1
	}
	var next int
	for i := 0; i < len(root.Children); i++ {
		cur := maxDepthForNary((*Node)(root.Children[i]))
		if cur > next {
			next = cur
		}
	}
	return 1 + next
}
