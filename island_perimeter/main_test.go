package main

import "testing"

func Test_islandPerimeter(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				grid: [][]int{
					{0, 1, 0, 0},
					{1, 1, 1, 0},
					{0, 1, 0, 0},
					{1, 1, 0, 0},
				},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := islandPerimeter(tt.args.grid); got != tt.want {
				t.Errorf("islandPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_islandPerimeter2(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				grid: [][]int{
					{0, 1, 0, 0},
					{1, 1, 1, 0},
					{0, 1, 0, 0},
					{1, 1, 0, 0},
				},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := islandPerimeter2(tt.args.grid); got != tt.want {
				t.Errorf("islandPerimeter2() = %v, want %v", got, tt.want)
			}
		})
	}
}
