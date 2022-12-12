package main

import (
	"bufio"
	"strings"
	"testing"
)

type Stuff struct {
	A string
	B int
}

func (s *Stuff) Set(value string) {
	s.A = value
}

func (s *Stuff) Incr() {
	s.B++
}

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
