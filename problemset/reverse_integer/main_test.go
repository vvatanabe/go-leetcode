package main

import "testing"

func Test_reverse(t *testing.T) {
	x := -54321
	want := -12345
	if got := reverse(x); got != want {
		t.Errorf("reverse() = %v, want %v", got, want)
	}
}

func Test_reverse2(t *testing.T) {
	x := -54321
	want := -12345
	if got := reverse2(x); got != want {
		t.Errorf("reverse2() = %v, want %v", got, want)
	}
}
