package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums := []int{4, 8, 2, 9, 0, 3, 7, 5, 6, 1}
	target := 10
	want := []int{1, 2}
	if got := twoSum(nums, target); !reflect.DeepEqual(got, want) {
		t.Errorf("twoSum() = %v, want %v", got, want)
	}
}
