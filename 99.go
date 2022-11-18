package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {
	arr := []*TreeNode{}
	pre := &TreeNode{}
	in(root, new(bool), &pre, &arr)
	if l := len(arr); l >= 2 {
		arr[0].Val, arr[l-1].Val = arr[l-1].Val, arr[0].Val
	}
}

func in(root *TreeNode, start *bool, pre **TreeNode, arr *[]*TreeNode) {
	if root == nil {
		return
	}
	in(root.Left, start, pre, arr)
	if *start && root.Val < (*pre).Val {
		*arr = append(*arr, *pre)
		*arr = append(*arr, root)
	}
	*pre = root
	if *start == false {
		*start = true
	}
	in(root.Right, start, pre, arr)
}
