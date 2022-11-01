package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	lower, higher := &ListNode{}, &ListNode{}
	l, h := lower, higher
	for p := head; p != nil; p = p.Next {
		if p.Val < x {
			l.Next = p
			l = p
		} else {
			h.Next = p
			h = p
		}
	}
	l.Next = higher.Next
	h.Next = nil
	return lower.Next
}
