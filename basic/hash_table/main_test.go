package main

import (
	"testing"
)

func Test_HashMap(t *testing.T) {
	m := NewHashMap(16)

	m.Put("aaa", "1")
	m.Put("bbb", "2")
	m.Put("ccc", "3")

	if got := m.Get("aaa"); got != "1" {
		t.Errorf("m.Get(\"aaa\") = %v, want %v", got, "1")
	}
	if got := m.Get("bbb"); got != "2" {
		t.Errorf("m.Get(\"bbb\") = %v, want %v", got, "2")
	}
	if got := m.Get("ccc"); got != "3" {
		t.Errorf("m.Get(\"ccc\") = %v, want %v", got, "3")
	}

	m.Put("aaa", "4")

	if got := m.Get("aaa"); got != "4" {
		t.Errorf("m.Get(\"aaa\") = %v, want %v", got, "4")
	}

	if got := m.Len(); got != 3 {
		t.Errorf("m.Len() = %v, want %v", got, 3)
	}
}
