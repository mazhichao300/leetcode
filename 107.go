package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		arr := []int{}
		q := []*TreeNode{}
		for _, p := range queue {
			arr = append(arr, p.Val)
			if p.Left != nil {
				q = append(q, p.Left)
			}
			if p.Right != nil {
				q = append(q, p.Right)
			}
		}
		ans = append([][]int{arr}, ans...)
		queue = q
	}
	return ans
}
