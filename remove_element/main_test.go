package main

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	want := 5
	if got := removeElement(nums, val); got != want {
		t.Errorf("removeElement() = %v, want %v", got, want)
	}
	want2 := []int{0, 1, 3, 0, 4}
	if !reflect.DeepEqual(nums[:want], want2) {
		t.Errorf("removeElement() = %v, want %v", nums, want2)
	}
}
