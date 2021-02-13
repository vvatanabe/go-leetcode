package main

import (
	"fmt"
)

func main() {
	head := &Node{
		Value: "apple",
		Next: &Node{
			Value: "banana",
			Next: &Node{
				Value: "cherry",
				Next: &Node{
					Value: "lemon",
					Next: &Node{
						Value: "peach",
						Next:  nil,
					},
				},
			},
		},
	}
	list := LinkedList{
		head: head,
		len:  5,
	}

	list.Print()
	list.Insert("raspberry")
	list.Print()
	list.DeleteAt(0)
	list.Print()
	list.InsertAt(0, "orange")
	list.Print()
	list.InsertAt(3, "apricot")
	list.Print()
}

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Size() int {
	return l.len
}

func (l *LinkedList) Insert(value string) {
	node := &Node{Value: value}
	if l.len == 0 {
		l.head = node
		l.len++
		return
	}
	head := l.head
	for i := 0; i < l.len; i++ {
		if head.Next == nil {
			head.Next = node
			l.len++
			return
		}
		head = head.Next
	}
}

func (l *LinkedList) InsertAt(position int, value string) {
	if position < 0 || position > l.len {
		return
	}
	node := &Node{Value: value}
	if position == 0 {
		if l.len > 0 {
			node.Next = l.head
		}
		l.head = node
		l.len++
		return
	}
	prev := l.GetAt(position - 1)
	prev.Next, node.Next = node, l.GetAt(position)
	l.len++
}

func (l *LinkedList) GetAt(position int) *Node {
	head := l.head
	if position < 0 || position > l.len-1 {
		return nil
	}
	for i := 0; i < position; i++ {
		head = head.Next
	}
	return head
}

func (l *LinkedList) DeleteAt(position int) {
	if l.len == 0 || position < 0 {
		return
	}
	if position == 0 {
		l.head = l.head.Next
		return
	}
	prev := l.GetAt(position - 1)
	if prev == nil {
		return
	}
	prev.Next = l.GetAt(position).Next
	l.len--
}

func (l *LinkedList) Print() {
	head := l.head
	for head != nil {
		fmt.Printf("%s -> ", head.Value)
		head = head.Next
	}
	fmt.Println()
}
