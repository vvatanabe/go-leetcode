package main

import "testing"

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n   int
		bad int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n:   5,
				bad: 4,
			},
			want: 4,
		},
		{
			args: args{
				n:   1,
				bad: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bad = tt.want
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
