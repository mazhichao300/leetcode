package main

import "fmt"

//Dfinition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	queue := []*ListNode{}

	for p := head.Next; p != nil; p = p.Next {
		queue = append(queue, p)
	}
	idx1 := 0
	idx2 := len(queue) - 1
	p := head
	direction := true
	for idx1 <= idx2 {
		if direction {
			p.Next = queue[idx2]
			idx2--
		} else {
			p.Next = queue[idx1]
			idx1++
		}
		p = p.Next

		direction = !direction
	}
	p.Next = nil
}

func print(p *ListNode) {
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
	fmt.Println()
}

func main() {
	p1 := &ListNode{Val: 1}
	p2 := &ListNode{Val: 2}
	p3 := &ListNode{Val: 3}
	p4 := &ListNode{Val: 4}
	p1.Next = p2
	p2.Next = p3
	p3.Next = p4
	reorderList(p1)
	print(p1)
}
