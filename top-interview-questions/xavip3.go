package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var ans *ListNode
	for head != nil {
		next := head.Next
		head.Next = ans
		ans = head
		head = next
	}
	return ans
}

func reverseListR(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	next := head.Next
}
