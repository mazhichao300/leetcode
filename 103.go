package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	direction := true
	for len(queue) > 0 {
		q := []*TreeNode{}
		arr := []int{}
		for _, p := range queue {
			if p.Left != nil {
				q = append(q, p.Left)
			}
			if p.Right != nil {
				q = append(q, p.Right)
			}
			if direction {
				arr = append(arr, p.Val)
			} else {
				arr = append([]int{p.Val}, arr...)
			}
		}
		ans = append(ans, arr)
		direction = !direction
		queue = q
	}
	return ans
}
