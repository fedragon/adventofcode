package main

import (
	"bufio"
	"github.com/fedragon/adventofcode/common"
	"os"
	"strings"
	"testing"
)

func TestPart1Solver_Solve(t *testing.T) {
	example := `
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`

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
			args{
				scanner: bufio.NewScanner(strings.NewReader(example)),
			},
			31,
			false,
		},
		{
			"solves part 1 with real input",
			args{
				scanner: bufio.NewScanner(common.Must(os.Open("../../data/day12"))),
			},
			361,
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

func TestPart2Solver_Solve(t *testing.T) {
	example := `
Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`

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
			"solves example in part 2",
			args{
				scanner: bufio.NewScanner(strings.NewReader(example)),
			},
			29,
			false,
		},
		{
			"solves part 2 with real input",
			args{
				scanner: bufio.NewScanner(common.Must(os.Open("../../data/day12"))),
			},
			354,
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
