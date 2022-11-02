package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {

}

func quick(start, end *ListNode, x int) {

}

func partition(start, end *ListNode, x int) *ListNode {
	p, q := start, end
	for q != nil {
		for p != nil && p.Val < x {
			p = p.Next
		}
		if p == nil {
			break
		}
		q = p.Next
		for q != nil && q.Val >= x {
			q = q.Next
		}
		if q != nil && p != nil {
			p.Val, q.Val = q.Val, p.Val
		}
	}
	return head
}
