package main

import "testing"

func Test_canWinNim(t *testing.T) {
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
				n: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canWinNim(tt.args.n); got != tt.want {
				t.Errorf("canWinNim() = %v, want %v", got, tt.want)
			}
		})
	}
}
