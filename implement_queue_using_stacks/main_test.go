package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if got := s.Pop(); got != 1 {
		t.Errorf("Pop() = %v, want %v", got, 1)
	}
	if got := s.Pop(); got != 2 {
		t.Errorf("Pop() = %v, want %v", got, 2)
	}
	if got := s.Pop(); got != 3 {
		t.Errorf("Pop() = %v, want %v", got, 3)
	}
	if got := s.Empty(); got != true {
		t.Errorf("Empty() = %v, want %v", got, true)
	}
}
