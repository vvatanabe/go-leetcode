package main

import "testing"

func Test_romanToInt(t *testing.T) {
	s := "IV"
	want := 4
	if got := romanToInt(s); got != want {
		t.Errorf("romanToInt() = %v, want %v", got, want)
	}
}
