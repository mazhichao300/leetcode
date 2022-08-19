package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	odd := true
	left, right, p, q := head, head.Next, head, head.Next

	for head != nil {
		next := head.Next
		if odd {
			p.Next = head
			p = head
		} else {
			q.Next = head
			q = head
		}
		head.Next = nil
		odd = !odd
		head = next
	}
	p.Next = right

	return left
}
