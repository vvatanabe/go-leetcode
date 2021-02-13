package main

import "testing"

func Test_findTheDifference(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			args: args{
				s: "abcd",
				t: "abcde",
			},
			want: byte('e'),
		},
		{
			args: args{
				s: "",
				t: "y",
			},
			want: byte('y'),
		},
		{
			args: args{
				s: "a",
				t: "aa",
			},
			want: byte('a'),
		},
		{
			args: args{
				s: "ae",
				t: "aea",
			},
			want: byte('a'),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDifference(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("findTheDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
