package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	tempA, tempB := headA, headB
	var cnt int
	for tempA != nil && tempB != nil {
		cnt++
		tempA = tempA.Next
		tempB = tempB.Next
	}
	lenA, lenB := cnt, cnt
	for tempA != nil {
		lenA++
		tempA = tempA.Next
	}
	for tempB != nil {
		lenB++
		tempB = tempB.Next
	}

	if lenA > lenB {
		n := lenA - lenB
		for n > 0 {
			n--
			headA = headA.Next
		}
	}
	if lenB > lenA {
		n := lenB - lenA
		for n > 0 {
			n--
			headB = headB.Next
		}
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	mp := make(map[*ListNode]struct{})
	for headA != nil {
		mp[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := mp[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
