package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1->2->3->4->5->NULL
// 1->NULL
// 2->1->NULL
// 3->2->1->NULL
// 4->3->2->1->NULL
// 5->4->3->2->1->NULL
func reverseList(head *ListNode) *ListNode {
	var node *ListNode
	for head != nil {
		node = &ListNode{
			Val:  head.Val,
			Next: node,
		}
		head = head.Next
	}
	return node
}
