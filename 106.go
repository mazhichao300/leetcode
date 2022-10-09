package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	num := len(postorder)
	if num == 0 {
		return nil
	}
	val := postorder[num-1]
	idx := 0
	for idx < num {
		if inorder[idx] == val {
			break
		}
		idx++
	}
	root := &TreeNode{Val: val}
	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:num-1])
	return root
}
