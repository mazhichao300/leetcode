/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	in(root, &ans)
	return ans
}

func in(p *TreeNode, ans *[]int) {
	if p == nil {
		return
	}
	in(p.Left, ans)
	*ans = append(*ans, p.Val)
	in(p.Right, ans)
}