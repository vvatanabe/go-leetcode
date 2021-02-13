package main

import "testing"

func Test_LinkedList(t *testing.T) {
	var list LinkedList
	list.Insert("apple")
	list.Insert("banana")
	list.Insert("cherry")

	if got := list.Size(); got != 3 {
		t.Errorf("list.Size() = %v, want %v", got, 3)
	}
	if got := list.GetAt(0).Value; got != "apple" {
		t.Errorf("list.GetAt(0).Value = %v, want %v", got, "apple")
	}
	if got := list.GetAt(1).Value; got != "banana" {
		t.Errorf("list.GetAt(1).Value = %v, want %v", got, "banana")
	}
	if got := list.GetAt(2).Value; got != "cherry" {
		t.Errorf("list.GetAt(2).Value = %v, want %v", got, "cherry")
	}

	list.InsertAt(0, "lemon")
	list.InsertAt(2, "peach")

	if got := list.GetAt(0).Value; got != "lemon" {
		t.Errorf("list.GetAt(0).Value = %v, want %v", got, "lemon")
	}
	if got := list.GetAt(2).Value; got != "peach" {
		t.Errorf("list.GetAt(2).Value = %v, want %v", got, "peach")
	}

	list.DeleteAt(1)
	if got := list.GetAt(1).Value; got != "peach" {
		t.Errorf("list.GetAt(1).Value = %v, want %v", got, "peach")
	}
	list.DeleteAt(3)
	if got := list.GetAt(2).Value; got != "banana" {
		t.Errorf("list.GetAt(3).Value = %v, want %v", got, "banana")
	}
}
