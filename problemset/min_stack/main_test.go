package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := Constructor()
	s.GetMin()
	s.Push(2)
	s.Push(0)
	s.Push(3)
	s.Push(0)
	if got := s.GetMin(); got != 0 {
		t.Errorf("GetMin() = %v, want %v", got, 0)
	}
	s.Pop()
	if got := s.GetMin(); got != 0 {
		t.Errorf("GetMin() = %v, want %v", got, 0)
	}
	s.Pop()
	if got := s.GetMin(); got != 0 {
		t.Errorf("GetMin() = %v, want %v", got, 0)
	}
	s.Pop()
	if got := s.GetMin(); got != 2 {
		t.Errorf("GetMin() = %v, want %v", got, 2)
	}
}
