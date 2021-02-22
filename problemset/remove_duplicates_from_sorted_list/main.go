package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	v := &ListNode{
		Next: head,
	}
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			*head = *head.Next
			continue
		}
		head = head.Next
	}
	return v.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
