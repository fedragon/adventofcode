package main

import "testing"

var grid = Grid{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func Test_isVisible(t *testing.T) {
	type args struct {
		grid Grid
		row  int
		col  int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"top-left 5 = true",
			args{
				grid: grid,
				row:  1,
				col:  1,
			},
			true,
		},
		{
			"top-middle 5 = true",
			args{
				grid: grid,
				row:  1,
				col:  2,
			},
			true,
		},
		{
			"top-right 1 = false",
			args{
				grid: grid,
				row:  1,
				col:  3,
			},
			false,
		},
		{
			"left-middle 5 = true",
			args{
				grid: grid,
				row:  2,
				col:  1,
			},
			true,
		},
		{
			"center 3 = false",
			args{
				grid: grid,
				row:  2,
				col:  2,
			},
			false,
		},
		{
			"right-middle 3 = true",
			args{
				grid: grid,
				row:  2,
				col:  3,
			},
			true,
		},
		{
			"bottom-middle 5 = true",
			args{
				grid: grid,
				row:  3,
				col:  2,
			},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isVisible(tt.args.grid, tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("isVisible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scenicScoreOf(t *testing.T) {
	type args struct {
		grid Grid
		row  int
		col  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"(3, 2) = 8",
			args{
				grid: grid,
				row:  3,
				col:  2,
			},
			8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scenicScoreOf(tt.args.grid, tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("scenicScoreOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
