package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := 1+minDepth(root.Left), 1+minDepth(root.Right)
	if l == 1 {
		return r
	}
	if r == 1 {
		return l
	}
	if l < r {
		return l
	}
	return r
}
