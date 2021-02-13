package main

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				n: 4,
			},
			want: true,
		},
		{
			args: args{
				n: 16,
			},
			want: true,
		},
		{
			args: args{
				n: 64,
			},
			want: true,
		},
		{
			args: args{
				n: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
