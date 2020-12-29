package main

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				x: 16,
			},
			want: 4,
		},
		{
			args: args{
				x: 36,
			},
			want: 6,
		},
		{
			args: args{
				x: 35,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run("sqrtx", func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
