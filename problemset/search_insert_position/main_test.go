package main

import "testing"

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	want := 2
	if got := searchInsert(nums, target); got != want {
		t.Errorf("searchInsert() = %v, want %v", got, want)
	}
}
