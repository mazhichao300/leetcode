package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	pre(root, &ans)
	return ans
}

func pre(p *TreeNode, ans *[]int) {
	if p == nil {
		return
	}
	*ans = append(*ans, p.Val)
	pre(p.Left, ans)
	pre(p.Right, ans)
}
