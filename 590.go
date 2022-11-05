package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	ans := []int{}
	post(root, &ans)
	return ans
}

func post(root *Node, ans *[]int) {
	if root == nil {
		return
	}
	for _, p := range root.Children {
		post(p, ans)
	}
	*ans = append(*ans, root.Val)
}
