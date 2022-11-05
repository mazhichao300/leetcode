package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	ans := 0
	back(root, 0, &ans)
	return ans
}

func back(root *TreeNode, cur int, ans *int) {
	if root == nil {
		return
	}
	cur = cur*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*ans += cur
	}
	back(root.Left, cur, ans)
	back(root.Right, cur, ans)
}
