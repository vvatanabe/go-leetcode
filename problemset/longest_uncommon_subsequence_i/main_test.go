package main

import "testing"

func Test_findLUSlength(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a: "aaa",
				b: "bbb",
			},
			want: 3,
		},
		{
			args: args{
				a: "ab",
				b: "a",
			},
			want: 2,
		},
		{
			args: args{
				a: "aaa",
				b: "aaa",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLUSlength(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}
