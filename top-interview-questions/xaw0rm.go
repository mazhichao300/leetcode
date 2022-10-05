package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var p1 *ListNode
	for head != slow {
		next := head.Next
		if p1 == nil {
			p1 = head
		} else {
			head.Next = p1
			p1 = head
		}
		head = next
	}
	// 表示链表长度是奇数
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if p1.Val != slow.Val {
			return false
		}
		slow = slow.Next
		p1 = p1.Next
	}
	return true
}
