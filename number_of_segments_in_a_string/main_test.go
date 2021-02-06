package main

import "testing"

func Test_countSegments(t *testing.T) {
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
				s: "Hello, my name is John",
			},
			want: 5,
		},
		{
			args: args{
				s: "                ",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSegments(tt.args.s); got != tt.want {
				t.Errorf("countSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countSegments2(t *testing.T) {
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
				s: "Hello, my name is John",
			},
			want: 5,
		},
		{
			args: args{
				s: "                ",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSegments2(tt.args.s); got != tt.want {
				t.Errorf("countSegments2() = %v, want %v", got, tt.want)
			}
		})
	}
}
