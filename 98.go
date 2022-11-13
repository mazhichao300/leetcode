package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	arr := []int{}
	return inOrder(root, &arr)
}

func inOrder(root *TreeNode, ans *[]int) bool {
	if root == nil {
		return true
	}
	if inOrder(root.Left, ans) == false {
		return false
	}
	if len(*ans) > 0 && root.Val < (*ans)[len(*ans)-1] {
		return false
	}
	*ans = append(*ans, root.Val)
	if inOrder(root.Right, ans) == false {
		return false
	}
	return true
}
