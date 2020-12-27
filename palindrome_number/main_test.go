package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	x := 11223344
	want := false
	if got := isPalindrome(x); got != want {
		t.Errorf("isPalindrome() = %v, want %v", got, want)
	}
}

func Test_isPalindrome2(t *testing.T) {
	x := 11223344
	want := false
	if got := isPalindrome2(x); got != want {
		t.Errorf("isPalindrome2() = %v, want %v", got, want)
	}
}
