package main

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var list []int

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	list = make([]int, 0)
	preOrder(root1)
	list1 := list
	list = make([]int, 0)
	preOrder(root2)
	list2 := list
	ans := make([]int, len(list1)+len(list2))
	idx := 0
	i := 0
	j := 0
	for i < len(list1) || j < len(list2) {
		if i >= len(list1) {
			ans[idx] = list2[j]
			j++
		} else if j >= len(list2) {
			ans[idx] = list1[i]
			i++
		} else if list1[i] < list2[j] {
			ans[idx] = list1[i]
			i++
		} else {
			ans[idx] = list2[j]
			j++
		}
		idx++
	}
	return ans
}

// func preOrderCmp(root1, root2 *TreeNode) {
// 	if root1 == nil {
// 		return
// 	}
// 	preOrder(root.Left)
// 	list = append(list, root.Val)
// 	preOrder(root.Right)
// }

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	preOrder(root.Left)
	list = append(list, root.Val)
	preOrder(root.Right)
}

func main() {

}
