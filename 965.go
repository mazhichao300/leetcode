package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return order(root, root.Val)
}

func order(root *TreeNode, pre int) bool {
	if root == nil {
		return true
	}
	if root.Val != pre {
		return false
	}
	if order(root.Left, pre) == false {
		return false
	}
	if order(root.Right, pre) == false {
		return false
	}
	return true
}
