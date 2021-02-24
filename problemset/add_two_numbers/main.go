package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		head  ListNode
		carry int
	)
	cur := &head
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return head.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	return helper(l1, l2, 0)
}

func helper(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	sum := carry
	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}
	return &ListNode{
		Val:  sum % 10,
		Next: helper(l1, l2, sum/10),
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
