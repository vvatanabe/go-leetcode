package main

import "testing"

func Test_maximumProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1, 2, 3",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "1, 2, 3, 4",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 24,
		},
		{
			name: "-1, -2, -3",
			args: args{
				nums: []int{-1, -2, -3},
			},
			want: -6,
		},
		{
			name: "-100,-98,-1,2,3,4",
			args: args{
				nums: []int{-100, -98, -1, 2, 3, 4},
			},
			want: 39200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
