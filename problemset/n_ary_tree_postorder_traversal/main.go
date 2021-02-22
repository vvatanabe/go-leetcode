package main

func postorder(root *Node) []int {
	var out []int
	if root != nil {
		for _, node := range root.Children {
			out = append(out, postorder(node)...)
		}
		out = append(out, root.Val)
	}
	return out
}

type Node struct {
	Val      int
	Children []*Node
}
