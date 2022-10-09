package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	first := true
	pre, p, q := head, head, head
	for q != nil && q.Next != nil {
		if first {
			first = false
		} else {
			pre = pre.Next
		}
		p = p.Next
		q = q.Next.Next
	}
	left := head
	if p == head {
		left = nil
	}
	right := p.Next
	pre.Next = nil
	root := &TreeNode{Val: p.Val}
	root.Left = sortedListToBST(left)
	root.Right = sortedListToBST(right)
	return root
}
