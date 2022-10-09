package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// 前序的首元素就是根节点
	val := preorder[0]
	root := &TreeNode{
		Val: val,
	}
	// 根据根节点划分左右子树
	idx := 0
	for idx < len(inorder) {
		if inorder[idx] == val {
			break
		}
		idx++
	}
	leftNum := idx
	root.Left = buildTree(preorder[1:1+leftNum], inorder[0:idx])
	root.Right = buildTree(preorder[1+leftNum:], inorder[idx+1:])

	return root
}
