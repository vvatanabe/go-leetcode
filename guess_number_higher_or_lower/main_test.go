package main

import "testing"

func Test_guessNumber(t *testing.T) {
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
				n: 10,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value = tt.want
			if got := guessNumber(tt.args.n); got != tt.want {
				t.Errorf("guessNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
