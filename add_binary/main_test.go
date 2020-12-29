package main

import "testing"

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				a: "1111",
				b: "1111",
			},
			want: "11110",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
