package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Input:
//
//   1
//  / \
// 2   3
//  \
//   5
//
// Output: ["1->2->5", "1->3"]
//
// Explanation: All root-to-leaf paths are: 1->2->5, 1->3
func binaryTreePaths(root *TreeNode) []string {
	var paths []string
	if root == nil {
		return paths
	}
	helper(root, "", &paths)
	return paths
}

func helper(r *TreeNode, path string, paths *[]string) {
	path += fmt.Sprintf("%d->", r.Val)
	if r.Left != nil {
		helper(r.Left, path, paths)
	}
	if r.Right != nil {
		helper(r.Right, path, paths)
	}
	if r.Left == nil && r.Right == nil {
		*paths = append(*paths, strings.TrimSuffix(path, "->"))
	}
}
