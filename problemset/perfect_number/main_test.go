package main

import "testing"

func Test_checkPerfectNumber(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				num: 496,
			},
			want: true,
		},
		{
			args: args{
				num: 8128,
			},
			want: true,
		},
		{
			args: args{
				num: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPerfectNumber(tt.args.num); got != tt.want {
				t.Errorf("checkPerfectNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPerfectNumber2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				num: 496,
			},
			want: true,
		},
		{
			args: args{
				num: 8128,
			},
			want: true,
		},
		{
			args: args{
				num: 2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPerfectNumber2(tt.args.num); got != tt.want {
				t.Errorf("checkPerfectNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}
