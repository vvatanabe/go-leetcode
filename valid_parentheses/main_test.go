package main

import "testing"

func Test_isValid(t *testing.T) {
	s := "()[]{}"
	want := true
	if got := isValid(s); got != want {
		t.Errorf("isValid() = %v, want %v", got, want)
	}
}

func Test_isValid2(t *testing.T) {
	s := "()[]{}"
	want := true
	if got := isValid2(s); got != want {
		t.Errorf("isValid2() = %v, want %v", got, want)
	}
}
