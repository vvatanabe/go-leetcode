package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	depth := 1
	for _, v := range root.Children {
		depth = max(depth, maxDepth(v)+1)
	}
	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Node struct {
	Val      int
	Children []*Node
}
