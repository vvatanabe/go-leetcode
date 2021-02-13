package main

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	digits := []int{4, 3, 2, 1}
	want := []int{4, 3, 2, 2}
	if got := plusOne(digits); !reflect.DeepEqual(got, want) {
		t.Errorf("plusOne() = %v, want %v", got, want)
	}
}
