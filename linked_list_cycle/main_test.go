package main

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "hasCycle",
			args: args{
				// [3,2,0,-4]
				head: func() *ListNode {
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
					return list
				}(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
