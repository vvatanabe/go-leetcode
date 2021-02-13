package main

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			args: args{
				n: 4,
			},
			want: 5,
		},
		{
			args: args{
				n: 5,
			},
			want: 8,
		},
		{
			args: args{
				n: 6,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run("climb stairs", func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
