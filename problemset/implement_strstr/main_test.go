package main

import "testing"

func Test_strStr(t *testing.T) {
	haystack := "mississippi"
	needle := "issip"
	want := 4
	if got := strStr(haystack, needle); got != want {
		t.Errorf("strStr() = %v, want %v", got, want)
	}
}
