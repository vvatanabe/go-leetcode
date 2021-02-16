package main

import "testing"

func Test_detectCapitalUse(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				word: "USA",
			},
			want: true,
		},
		{
			args: args{
				word: "leetcode",
			},
			want: true,
		},
		{
			args: args{
				word: "Google",
			},
			want: true,
		},
		{
			args: args{
				word: "FlaG",
			},
			want: false,
		},
		{
			args: args{
				word: "ffffffffffffffffffffF",
			},
			want: false,
		},
		{
			args: args{
				word: "mL",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCapitalUse(tt.args.word); got != tt.want {
				t.Errorf("detectCapitalUse() = %v, want %v", got, tt.want)
			}
		})
	}
}
