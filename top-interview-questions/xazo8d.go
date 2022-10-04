package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	list := []int{}
	dfs(root, k, &list)
	return list[k-1]
}

func dfs(root *TreeNode, k int, list *[]int) {
	if root == nil {
		return
	}
	if len(*list) >= k {
		return
	}
	if root.Left != nil {
		dfs(root.Left, k, list)
	}
	*list = append(*list, root.Val)
	if root.Right != nil {
		dfs(root.Right, k, list)
	}
}
