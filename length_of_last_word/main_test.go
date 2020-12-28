package main

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	s := "Hello World"
	want := 5
	if got := lengthOfLastWord(s); got != want {
		t.Errorf("lengthOfLastWord() = %v, want %v", got, want)
	}
}
