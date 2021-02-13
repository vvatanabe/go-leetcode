package main

import "fmt"

func main() {
	m := NewHashMap(16)
	m.Put("aaa", "1")
	m.Put("bbb", "2")
	m.Put("ccc", "3")
	m.Put("aaa", "4")
	fmt.Println(m.Get("aaa"))
	fmt.Println(m.Get("bbb"))
	fmt.Println(m.Get("ccc"))
	fmt.Println(m.Len())
}

type Node struct {
	key   string
	value string
	hash  int
	next  *Node
}

func NewHashMap(capacity int) *HashMap {
	return &HashMap{
		table:    make([]*Node, capacity),
		capacity: capacity,
	}
}

type HashMap struct {
	table    []*Node
	size     int
	capacity int
}

func (m *HashMap) Put(k, v string) {
	hash := m.hashCode(k)
	n := m.table[hash]
	for {
		if n == nil {
			break
		}
		if n.key == k {
			n.value = v
			return
		}
		n = n.next
	}
	m.table[hash] = &Node{
		key:   k,
		value: v,
		hash:  hash,
		next:  m.table[hash],
	}
	m.size++
}

func (m *HashMap) Get(k string) string {
	if k == "" {
		return ""
	}
	hash := m.hashCode(k)
	n := m.table[hash]
	for {
		if n == nil {
			break
		}
		if n.key == k {
			return n.value
		}
		n = n.next
	}
	return ""
}

func (m *HashMap) Len() int {
	return m.size
}

func (m *HashMap) hashCode(k string) int {
	var num int
	for i := 0; i < len(k); i++ {
		num += int(k[i])
	}
	return num % m.capacity
}
