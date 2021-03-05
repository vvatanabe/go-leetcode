package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	var levels []float64
	if root == nil {
		return levels
	}
	var queue Queue
	queue.Push(root)
	for queue.Len() != 0 {
		var sum, cnt float64
		size := queue.Len()
		for i := 0; i < size; i++ {
			cur := queue.Pop()
			sum += float64(cur.Val)
			cnt++
			if cur.Left != nil {
				queue.Push(cur.Left)
			}
			if cur.Right != nil {
				queue.Push(cur.Right)
			}
		}
		levels = append(levels, sum/cnt)
	}
	return levels
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	queue []*TreeNode
	size  int
}

func (q *Queue) Len() int {
	return q.size
}

func (q *Queue) Push(node *TreeNode) {
	q.queue = append(q.queue, node)
	q.size++
}

func (q *Queue) Pop() *TreeNode {
	if q.size == 0 {
		return nil
	}
	cur := q.queue[0]
	q.queue = q.queue[1:]
	q.size--
	return cur
}
