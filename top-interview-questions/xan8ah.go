package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lena, lenb := 0, 0
	p, q := headA, headB
	for p != nil {
		lena++
		p = p.Next
	}
	for q != nil {
		lenb++
		q = q.Next
	}
	if lena > lenb {
		lena -= lenb
		lenb = 0
	} else {
		lenb -= lena
		lena = 0
	}
	p, q = headA, headB
	for p != nil && q != nil {
		if p == q {
			return p
		}
		if lena > 0 {
			p = p.Next
			lena--
		} else if lenb > 0 {
			q = q.Next
			lenb--
		} else {
			p = p.Next
			q = q.Next
		}
	}
	return nil
}
