package main

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {

	// args: [3,2,0,-4]
	// want: 1

	list := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}
	list.Next.Next.Next.Next = list.Next

	type args struct {
		head *ListNode
	}
	test := struct {
		args args
		want *ListNode
	}{
		args: args{
			head: list,
		},
		want: list.Next,
	}

	if got := detectCycle(test.args.head); !reflect.DeepEqual(got, test.want) {
		t.Errorf("detectCycle() = %v, want %v", got, test.want)
	}
}
