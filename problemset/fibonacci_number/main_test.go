package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 0,
			},
			want: 0,
		},
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
			want: 1,
		},
		{
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			args: args{
				n: 4,
			},
			want: 3,
		},
		{
			args: args{
				n: 30,
			},
			want: 832040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
