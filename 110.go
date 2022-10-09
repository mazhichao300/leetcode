package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return height(root) != -1
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := height(root.Left)
	right := height(root.Right)
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}

	if left < right {
		left = right
	}
	return left + 1
}
