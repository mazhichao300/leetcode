package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var pre *TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	//left, right := root.Left, root.Right
	if pre != nil {
		pre.Left = nil
		pre.Right = root
	}
	pre = root
	flatten(root.Left)
	flatten(root.Right)
}

func flatten1(root *TreeNode) {
	preOrder(root, nil)
}

func preOrder(root, pre *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if pre != nil {
		pre.Right = root
	}

	left, right := root.Left, root.Right
	root.Left = nil
	l := preOrder(left, root)
	var r *TreeNode
	if l != nil {
		r = preOrder(right, l)
	} else {
		r = preOrder(right, root)
	}
	if r != nil {
		return r
	}
	if l != nil {
		return l
	}
	return root
}
