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
	prev := -1
	min := math.MaxInt64
	helper(root, &prev, &min)
	return min
}

func helper(root *TreeNode, prev, min *int) {
	if root != nil {
		helper(root.Left, prev, min)
		if *prev != -1 {
			if root.Val-*prev < *min {
				*min = root.Val - *prev
			}
		}
		*prev = root.Val
		helper(root.Right, prev, min)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
