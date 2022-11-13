package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	m := map[int]int{}
	find(root, m)
	max := -1
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	ans := []int{}
	for i, v := range m {
		if v == max {
			ans = append(ans, i)
		}
	}
	return ans
}

func find(root *TreeNode, m map[int]int) {
	if root == nil {
		return
	}
	find(root.Left, m)
	m[root.Val]++
	find(root.Right, m)
}
