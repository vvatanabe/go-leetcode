package main

import "testing"

func Test_firstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "leetcode",
			},
			want: 0,
		},
		{
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
		{
			args: args{
				s: "cc",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstUniqChar2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "leetcode",
			},
			want: 0,
		},
		{
			args: args{
				s: "loveleetcode",
			},
			want: 2,
		},
		{
			args: args{
				s: "cc",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqChar2(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar2() = %v, want %v", got, tt.want)
			}
		})
	}
}
