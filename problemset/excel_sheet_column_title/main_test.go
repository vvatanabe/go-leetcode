package main

import "testing"

func Test_convertToTitle(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				n: 1,
			},
			want: "A",
		},
		{
			args: args{
				n: 2,
			},
			want: "B",
		},
		{
			args: args{
				n: 3,
			},
			want: "C",
		},
		{
			args: args{
				n: 26,
			},
			want: "Z",
		},
		{
			args: args{
				n: 27,
			},
			want: "AA",
		},
		{
			args: args{
				n: 28,
			},
			want: "AB",
		},
		{
			args: args{
				n: 701,
			},
			want: "ZY",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.n); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
