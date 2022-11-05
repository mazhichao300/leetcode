package main

import "fmt"

//* Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	root := &ListNode{0, head}
	quick(root, nil)
	return root.Next
}

func quick(start, end *ListNode) {
	p := partition(start, end)
	if p == nil {
		return
	}
	quick(start, p)
	quick(p, end)
}

func partition(start, end *ListNode) *ListNode {
	if start.Next == end {
		return nil
	}
	p := FindMid(start, end)
	x := p.Val
	low, high := &ListNode{}, &ListNode{}
	l, h := low, high
	for i := start.Next; i != end; i = i.Next {
		if i == p {
			continue
		}
		if i.Val < x {
			l.Next = i
			l = i
		} else {
			h.Next = i
			h = i
		}
	}
	l.Next = p
	h.Next = end
	p.Next = high.Next
	start.Next = low.Next
	return p
}

func FindMid(start, end *ListNode) *ListNode {
	f, s := start, start
	for s.Next != end {
		f = f.Next
		s = s.Next
		if s.Next == end {
			break
		}
		s = s.Next
	}
	return f
}

func PrintList(a *ListNode) {
	for a != nil {
		fmt.Print(a.Val, " ")
		a = a.Next
	}
	fmt.Println("")
}

func MergeList(a, b *ListNode) *ListNode {
	ans := a
	cur := ans
	p, q := a.Next, b.Next
	for p != nil || q != nil {
		if p == nil || q.Val < p.Val {
			cur.Next = q
			cur = q
			q = q.Next
		} else {
			cur.Next = p
			cur = p
			p = p.Next
		}
	}
	cur.Next = nil
	return ans
}

func main() {
	a := &ListNode{4, nil}
	b := &ListNode{2, nil}
	c := &ListNode{1, nil}
	d := &ListNode{3, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	a = sortList(a)
	PrintList(a)
	root := &ListNode{0, a}
	p := FindMid(root, a)
	fmt.Println(p.Val)
}
