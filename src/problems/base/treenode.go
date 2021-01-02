package base

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NarayTreeNode struct {
	Val      int
	Children []*NarayTreeNode
}
