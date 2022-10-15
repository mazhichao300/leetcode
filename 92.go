package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	var leftNodePre *ListNode
	var leftNode *ListNode
	p := head
	for idx := 1; idx <= right; idx++ {
		next := p.Next
		if idx == left {
			leftNode = p
		} else if idx == left-1 {
			leftNodePre = p
		} else if idx > left {
			if leftNodePre != nil {
				p.Next = leftNodePre.Next
				leftNodePre.Next = p
			} else {
				p.Next = head
				head = p
			}
		}

		if idx == right {
			leftNode.Next = next
		}

		p = next
	}
	return head
}
