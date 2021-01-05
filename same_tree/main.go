package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		return p.Val == q.Val &&
			isSameTree(p.Left, q.Left) &&
			isSameTree(p.Right, q.Right)
	}
	if p == nil && q == nil {
		return true
	}
	return false
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return true
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
