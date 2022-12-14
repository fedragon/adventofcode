package main

import (
	"bufio"
	"fmt"
	. "github.com/fedragon/adventofcode/common"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []Segment
	}{
		{
			"parses a path",
			args{line: "498,4 -> 498,6 -> 496,6"},
			[]Segment{
				{
					{
						X: 498,
						Y: 4,
					},
					{
						X: 498,
						Y: 6,
					},
				},
				{
					{
						X: 498,
						Y: 6,
					},
					{
						X: 496,
						Y: 6,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSegment_Walk(t *testing.T) {
	tests := []struct {
		name string
		s    Segment
		want []Point
	}{
		{
			"walks left",
			Segment{{X: 0, Y: 2}, {X: 0, Y: 0}},
			[]Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}},
		},
		{
			"walks right",
			Segment{{X: 0, Y: 0}, {X: 0, Y: 2}},
			[]Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}},
		},
		{
			"walks up",
			Segment{{X: 2, Y: 0}, {X: 0, Y: 0}},
			[]Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}},
		},
		{
			"walks down",
			Segment{{X: 0, Y: 0}, {X: 2, Y: 0}},
			[]Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Walk(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Walk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart1Solver_Solve(t *testing.T) {
	type args struct {
		scanner *bufio.Scanner
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"solves example in part 1",
			args{bufio.NewScanner(strings.NewReader(`498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`))},
			24,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Part1Solver{}
			got, err := ds.Solve(tt.args.scanner)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Solve() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrid_PourSand(t *testing.T) {
	type fields struct {
		grid Grid
	}
	tests := []struct {
		name      string
		fields    fields
		want      bool
		wantCoord Point
	}{
		{
			"pours sand vertically until tile is blocked",
			fields{Grid{
				Tiles: map[Point]Tile{
					Point{X: 500, Y: 0}: Source,
					Point{X: 498, Y: 3}: Rock, Point{X: 499, Y: 3}: Rock, Point{X: 500, Y: 3}: Rock, Point{X: 501, Y: 3}: Rock, Point{X: 502, Y: 3}: Rock,
				},
				MaxY: 3,
				MinX: 498,
				MaxX: 502,
			}},
			true,
			Point{Y: 2, X: 500},
		},
		{
			"pours sand diagonally to the left if tile below is blocked and the one to the left is not",
			fields{Grid{
				Tiles: map[Point]Tile{
					Point{X: 500, Y: 0}: Source,
					Point{X: 500, Y: 2}: Sand,
					Point{X: 498, Y: 3}: Rock, Point{X: 499, Y: 3}: Rock, Point{X: 500, Y: 3}: Rock, Point{X: 501, Y: 3}: Rock, Point{X: 502, Y: 3}: Rock,
				},
				MaxY: 3,
				MinX: 498,
				MaxX: 502,
			}},
			true,
			Point{Y: 2, X: 499},
		},
		{
			"pours sand diagonally to the right if both below & left are blocked and the one to the right is not",
			fields{Grid{
				Tiles: map[Point]Tile{
					Point{X: 500, Y: 0}: Source,
					Point{X: 499, Y: 2}: Sand, Point{X: 500, Y: 2}: Sand,
					Point{X: 498, Y: 3}: Rock, Point{X: 499, Y: 3}: Rock, Point{X: 500, Y: 3}: Rock, Point{X: 501, Y: 3}: Rock, Point{X: 502, Y: 3}: Rock,
				},
				MaxY: 3,
				MinX: 498,
				MaxX: 502,
			}},
			true,
			Point{Y: 2, X: 501},
		},
		{
			"pours sand into the abyss, when nothing blocks its vertical fall",
			fields{Grid{
				Tiles: map[Point]Tile{
					Point{X: 500, Y: 0}: Source,
				},
				MaxY: 2,
				MinX: 500,
				MaxX: 500,
			}},
			false,
			Point{},
		},
		{
			"pours sand into the abyss, when no tiles are available",
			fields{Grid{
				Tiles: map[Point]Tile{
					Point{X: 500, Y: 0}: Source,
					Point{X: 500, Y: 1}: Sand,
					Point{X: 499, Y: 2}: Sand, Point{X: 500, Y: 2}: Sand, Point{X: 501, Y: 2}: Sand,
					Point{X: 498, Y: 3}: Rock, Point{X: 499, Y: 3}: Rock, Point{X: 500, Y: 3}: Rock, Point{X: 501, Y: 3}: Rock, Point{X: 502, Y: 3}: Rock,
				},
				MaxY: 3,
				MinX: 498,
				MaxX: 502,
			}},
			false,
			Point{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := tt.fields.grid
			got := grid.PourSand(SourceX, SourceY+1)
			fmt.Println(grid.String())

			if assert.Equal(t, tt.want, got) {
				if tt.want {
					assert.Equal(t, Sand, grid.Tiles[tt.wantCoord])
				}
			}
		})
	}
}

func TestPart2Solver_Solve(t *testing.T) {
	type args struct {
		scanner *bufio.Scanner
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"solves example in part 1",
			args{bufio.NewScanner(strings.NewReader(`498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`))},
			93,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Part2Solver{}
			got, err := ds.Solve(tt.args.scanner)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Solve() got = %v, want %v", got, tt.want)
			}
		})
	}
}
