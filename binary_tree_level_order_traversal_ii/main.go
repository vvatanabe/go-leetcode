package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var v [][]int
	helper(root, 0, &v)
	return v
}

func helper(root *TreeNode, i int, v *[][]int) {
	if root == nil {
		return
	}
	if len(*v) < i+1 {
		*v = append([][]int{{}}, *v...)
	}
	helper(root.Left, i+1, v)
	helper(root.Right, i+1, v)
	(*v)[len(*v)-i-1] = append((*v)[len(*v)-i-1], root.Val)
}
