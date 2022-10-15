package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	dfs(root, targetSum, []int{}, &ans)
	return ans
}

func dfs(root *TreeNode, targetSum int, arr []int, ans *[][]int) {
	if root == nil {
		return
	}
	arr = append(arr, root.Val)
	l := len(arr)
	// 找到一个结果
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		*ans = append(*ans, append([]int{}, arr...))
	}
	dfs(root.Left, targetSum-root.Val, arr, ans)
	arr = arr[:l]
	dfs(root.Right, targetSum-root.Val, arr, ans)
	arr = arr[:l]
}
