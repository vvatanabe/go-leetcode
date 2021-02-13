package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	h := ListNode{Next: head}
	prev := &h
	current := h.Next
	for current != nil && current.Next != nil {
		c := current
		for c != nil && c.Next != nil && c.Val == c.Next.Val {
			c = c.Next
		}
		if c != current {
			prev.Next = c.Next
			current = c.Next
		} else {
			prev = current
			current = current.Next
		}
	}
	return h.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
