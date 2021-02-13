package main

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	want := 5
	if got := removeDuplicates(nums); got != want {
		t.Errorf("removeDuplicates() = %v, want %v", got, want)
	}
	want2 := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(nums[:5], want2) {
		t.Errorf("removeDuplicates() = %v, want %v", nums, want2)
	}
}
