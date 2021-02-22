package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func preorder(root *Node) []int {
	var out []int
	if root != nil {
		out = append(out, root.Val)
		for _, node := range root.Children {
			out = append(out, preorder(node)...)
		}
	}
	return out
}

type Node struct {
	Val      int
	Children []*Node
}
