package main

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				num: 81,
			},
			want: true,
		},
		{
			args: args{
				num: 64,
			},
			want: true,
		},
		{
			args: args{
				num: 49,
			},
			want: true,
		},
		{
			args: args{
				num: 49,
			},
			want: true,
		},
		{
			args: args{
				num: 14,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
