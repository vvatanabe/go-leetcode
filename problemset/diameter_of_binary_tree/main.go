package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 1
	helper(root, &diameter)
	return diameter - 1
}

func helper(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}
	l := helper(root.Left, diameter)
	r := helper(root.Right, diameter)
	*diameter = max(*diameter, l+r+1)
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
