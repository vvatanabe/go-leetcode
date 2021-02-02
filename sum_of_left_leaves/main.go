package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	return helper(root, false)
}

func helper(node *TreeNode, isLeft bool) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		if isLeft {
			return node.Val
		}
		return 0
	}
	return helper(node.Left, true) +
		helper(node.Right, false)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
