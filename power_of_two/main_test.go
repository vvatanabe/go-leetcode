package main

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
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
				n: 2,
			},
			want: true,
		},
		{
			args: args{
				n: 4,
			},
			want: true,
		},
		{
			args: args{
				n: 8,
			},
			want: true,
		},
		{
			args: args{
				n: 9,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPowerOfTwo2(t *testing.T) {
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
				n: 2,
			},
			want: true,
		},
		{
			args: args{
				n: 4,
			},
			want: true,
		},
		{
			args: args{
				n: 8,
			},
			want: true,
		},
		{
			args: args{
				n: 9,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo2(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwo2() = %v, want %v", got, tt.want)
			}
		})
	}
}
