package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]struct{})
	node := head
	for node != nil {
		_, ok := m[node]
		if ok {
			return true
		}
		m[node] = struct{}{}
		node = node.Next
	}
	return false
}
