package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	current := 0
	ans := 0
	rec(root, &current, &ans)
	return ans
}

func rec(root *TreeNode, current *int, ans *int) {
	if root == nil {
		return
	}
	*current = *current*2 + root.Val
	// 找到叶子
	if root.Left == nil && root.Right == nil {
		*ans += *current
	}
	if root.Left != nil {
		rec(root.Left, current, ans)
	}
	if root.Right != nil {
		rec(root.Right, current, ans)
	}
	// 回溯
	*current -= root.Val
	*current /= 2
}

func main() {

}
