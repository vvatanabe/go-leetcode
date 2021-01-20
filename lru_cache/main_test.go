package main

import "testing"

func TestLRUCache(t *testing.T) {
	s := Constructor(2)
	s.Put(1, 1)
	s.Put(2, 2)
	if got := s.Get(1); got != 1 {
		t.Errorf("Pop() = %v, want %v", got, 1)
	}
	s.Put(3, 3)
	if got := s.Get(2); got != -1 {
		t.Errorf("Pop() = %v, want %v", got, -1)
	}
	s.Put(4, 4)
	if got := s.Get(1); got != -1 {
		t.Errorf("Pop() = %v, want %v", got, -1)
	}
	if got := s.Get(3); got != 3 {
		t.Errorf("Pop() = %v, want %v", got, 3)
	}
	if got := s.Get(4); got != 4 {
		t.Errorf("Pop() = %v, want %v", got, 4)
	}
}
