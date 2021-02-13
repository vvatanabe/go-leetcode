package main

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{1, 3, 8, 18, 31, 43, 68, 72, 101},
				k:    68,
			},
			want: 6,
		},
		{
			args: args{
				nums: []int{1, 3, 8, 18, 31, 43, 68, 72, 101},
				k:    102,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
