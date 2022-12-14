package main

import (
	"bufio"
	"github.com/fedragon/adventofcode/common"
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
						Y: 498,
						X: 4,
					},
					{
						Y: 498,
						X: 6,
					},
				},
				{
					{
						Y: 498,
						X: 6,
					},
					{
						Y: 496,
						X: 6,
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
		want []common.Coord
	}{
		{
			"walks left",
			Segment{{X: 0, Y: 2}, {X: 0, Y: 0}},
			[]common.Coord{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}},
		},
		{
			"walks right",
			Segment{{X: 0, Y: 0}, {X: 0, Y: 2}},
			[]common.Coord{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}},
		},
		{
			"walks up",
			Segment{{X: 2, Y: 0}, {X: 0, Y: 0}},
			[]common.Coord{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}},
		},
		{
			"walks down",
			Segment{{X: 0, Y: 0}, {X: 2, Y: 0}},
			[]common.Coord{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}},
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
			0, // TODO implement
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
