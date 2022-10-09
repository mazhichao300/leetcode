package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDep, rightDep := 1, 1
	if root.Left != nil {
		leftDep += maxDepth(root.Left)
	}
	if root.Right != nil {
		rightDep += maxDepth(root.Right)
	}
	if leftDep > rightDep {
		return leftDep
	}
	return rightDep
}
