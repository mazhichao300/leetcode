package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := map[*Node]*Node{}
	q := []*Node{}
	for head != nil {
		tmp := &Node{head.Val, head.Next, head.Random}
		q = append(q, tmp)
		m[head] = tmp
		head = head.Next
	}
	for i := 0; i < len(q); i++ {
		if i+1 < len(q) {
			q[i].Next = q[i+1]
		}
		if q[i].Random != nil {
			q[i].Random = m[q[i].Random]
		}
	}
	return q[0]
}
