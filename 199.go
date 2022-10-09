package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	size := 1

	for size > 0 {
		p := queue[size-1]
		ans = append(ans, p.Val)
		q := []*TreeNode{}
		for _, n := range queue {
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		queue = q
		size = len(queue)
	}
	return ans
}
