package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	_, min := helper(root, -1, math.MaxInt32)
	return min
}

func helper(root *TreeNode, prev, min int) (int, int) {
	if root == nil {
		return prev, min
	}
	prev, min = helper(root.Left, prev, min)
	if prev != -1 {
		if diff := root.Val - prev; diff < min {
			min = diff
		}
	}
	return helper(root.Right, root.Val, min)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
