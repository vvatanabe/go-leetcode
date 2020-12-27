package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	want := "fl"
	if got := longestCommonPrefix(strs); got != want {
		t.Errorf("longestCommonPrefix() = %v, want %v", got, want)
	}
}
