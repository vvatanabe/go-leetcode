package main

import (
	"reflect"
	"testing"
)

func slice2List(slice ...int) *ListNode {
	var head ListNode
	cur := &head
	for _, v := range slice {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head.Next
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				l1: slice2List(2, 4, 3),
				l2: slice2List(5, 6, 4),
			},
			want: slice2List(7, 0, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addTwoNumber2(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				l1: slice2List(2, 4, 3),
				l2: slice2List(5, 6, 4),
			},
			want: slice2List(7, 0, 8),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers2(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers2() = %v, want %v", got, tt.want)
			}
		})
	}
}
