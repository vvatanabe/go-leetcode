package main

import "testing"

func Test_isPowerOfThree(t *testing.T) {
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
				n: 3,
			},
			want: true,
		},
		{
			args: args{
				n: 9,
			},
			want: true,
		},
		{
			args: args{
				n: 27,
			},
			want: true,
		},
		{
			args: args{
				n: 45,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPowerOfThree2(t *testing.T) {
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
				n: 3,
			},
			want: true,
		},
		{
			args: args{
				n: 9,
			},
			want: true,
		},
		{
			args: args{
				n: 27,
			},
			want: true,
		},
		{
			args: args{
				n: 45,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree2(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThree2() = %v, want %v", got, tt.want)
			}
		})
	}
}
