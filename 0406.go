package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 直接根据p的情况判断
// 如果p有右子树，则应该返回右子树的最左节点
// 如果没有右子树，则向根节点回溯，找第一个比p大的节点
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}

	//
	if p.Right != nil {
		p = p.Right
		for p.Left != nil {
			p = p.Left
		}
		return p
	}

	num := 0
	var ans *TreeNode

	for root != p {
		if root.Val > p.Val && (ans == nil || root.Val < num) {
			ans = root
			num = root.Val
		}
		if p.Val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return ans
}

// 中序遍历，返回第一个比p大的节点
func inorderSuccessor2(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}
	left := inorderSuccessor2(root.Left, p)
	if left != nil {
		return left
	}
	if root.Val > p.Val {
		return root
	}
	return inorderSuccessor2(root.Right, p)
}

// 先中序遍历得到数组，再取p的下一个节点
func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}
	ans := &[]*TreeNode{}
	inorder(root, ans)
	idx := 0
	for i, ptr := range *ans {
		if p == ptr {
			idx = i
			break
		}
	}
	if idx == len(*ans)-1 {
		return nil
	}
	return (*ans)[idx+1]
}

func inorder(root *TreeNode, ans *[]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorder(root.Left, ans)
	}
	*ans = append(*ans, root)
	if root.Right != nil {
		inorder(root.Right, ans)
	}
}

func main() {

}
