package main

func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	post(root, &ans)
	return ans
}

func post(p *TreeNode, ans *[]int) {
	if p == nil {
		return
	}
	post(p.Left, ans)
	post(p.Right, ans)
	*ans = append(*ans, p.Val)
}
