package main

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func TestInteger_CompareInteger(t *testing.T) {
	type args struct {
		v Value
	}
	tests := []struct {
		name string
		i    Int
		args args
		want Outcome
	}{
		{
			"returns -1 if a < b",
			Int(1),
			args{Int(2)},
			-1,
		},
		{
			"returns 0 if a == b",
			Int(1),
			args{Int(1)},
			0,
		},
		{
			"returns 1 if a > b",
			Int(2),
			args{Int(1)},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Compare(tt.args.v); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteger_CompareList(t *testing.T) {
	type args struct {
		v Value
	}
	tests := []struct {
		name string
		i    Int
		args args
		want Outcome
	}{
		{
			"returns -1 if len(b) == 1 and a[0] < b[0]",
			Int(1),
			args{List{Int(2)}},
			-1,
		},
		{
			"returns 0 if len(b) == 1 and a[0] == b[0]",
			Int(1),
			args{List{Int(1)}},
			0,
		},
		{
			"returns 1 if len(b) == 1 and a[0] > b[0]",
			Int(2),
			args{List{Int(1)}},
			1,
		},
		{
			"returns -1 if len(b) > 1 and a[0] < b[0]",
			Int(1),
			args{List{Int(2), Int(3)}},
			-1,
		},
		{
			"returns -1 if len(b) > 1 and a[0] == b[0]",
			Int(1),
			args{List{Int(1), Int(2)}},
			-1,
		},
		{
			"returns 1 if len(b) > 1 and a[0] > b[0]",
			Int(2),
			args{List{Int(1), Int(3)}},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Compare(tt.args.v); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_CompareList(t *testing.T) {
	type args struct {
		v Value
	}
	tests := []struct {
		name string
		i    List
		args args
		want Outcome
	}{
		{
			"returns -1 if len(a) == len(b) and a[0] < b[0]",
			List{Int(1)},
			args{List{Int(2)}},
			-1,
		},
		{
			"returns 0 if len(a) == len(b) and a[0] == b[0]",
			List{Int(1)},
			args{List{Int(1)}},
			0,
		},
		{
			"returns 1 if len(a) == len(b) and a[0] > b[0]",
			List{Int(2)},
			args{List{Int(1)}},
			1,
		},
		{
			"returns -1 if len(a) < len(b) and a[x] < b[x] for all items in a",
			List{Int(1)},
			args{List{Int(2), Int(3)}},
			-1,
		},
		{
			"returns -1 if len(a) < len(b) and a[x] == b[x] for all items in a",
			List{Int(1), Int(2)},
			args{List{Int(1), Int(2), Int(3)}},
			-1,
		},
		{
			"returns -1 if len(a) < len(b) and there exists at least one item in a for which a[x] < b[x]",
			List{Int(2), Int(2)},
			args{List{Int(2), Int(3), Int(5)}},
			-1,
		},
		{
			"returns 1 if len(a) < len(b) and there exists at least one item in a for which a[x] > b[x]",
			List{Int(2), Int(4)},
			args{List{Int(2), Int(3), Int(5)}},
			1,
		},
		{
			"returns 1 if len(a) > len(b) and a[x] == b[x] for all items is a",
			List{Int(1), Int(2), Int(3)},
			args{List{Int(1), Int(2)}},
			1,
		},
		{
			"returns 1 if len(a) > len(b) and a[x] == b[x] for all items is a",
			List{Int(1), Int(2), Int(3)},
			args{List{Int(1), Int(2)}},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Compare(tt.args.v); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_CompareInteger(t *testing.T) {
	type args struct {
		v Value
	}
	tests := []struct {
		name string
		i    List
		args args
		want Outcome
	}{
		{
			"returns -1 if b is integer and a[0] < b",
			List{Int(1)},
			args{Int(2)},
			-1,
		},
		{
			"returns 0 if b is integer and a[0] == b",
			List{Int(1)},
			args{Int(1)},
			0,
		},
		{
			"returns 1 if b is integer and a[0] > b",
			List{Int(2)},
			args{Int(1)},
			1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Compare(tt.args.v); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Sample(t *testing.T) {
	type args struct {
		v Value
	}
	tests := []struct {
		name string
		i    List
		args args
		want Outcome
	}{
		// [[],[1,[3,[0]]],[]]
		// [[],[[0,[5,3,0,1,0],[3,0,5,7],10,[2,8,5,0]],10,[2,4,[1],[5,6,7],[]],[]]]

		// [[8,[],[10]],[]]
		// [[[[8,5,6,6,5],1,[10]],[]],[],[],[7],[2,2]]
		{
			"returns -1 if b is integer and a[0] < b",
			List{List{Int(8), List{}, List{Int(10)}, List{}}},
			args{List{List{List{List{Int(8), Int(5), Int(6), Int(6), Int(5)}, Int(1), List{Int(10)}}, List{}}, List{}, List{}, List{Int(7)}, List{Int(2), Int(2)}}},
			-1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Compare(tt.args.v); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want List
	}{
		{
			"parses an empty list",
			args{line: "[]"},
			nil,
		},
		{
			"parses a list with a single integer",
			args{line: "[99]"},
			List{Int(99)},
		},
		{
			"parses a list of lists",
			args{line: "[[11],[99]]"},
			List{List{Int(11)}, List{Int(99)}},
		},
		{
			"parses a list of mixed elements",
			args{line: "[11,[99]]"},
			List{Int(11), List{Int(99)}},
		},
		{
			"parses another list of mixed elements",
			args{line: "[[11],22,[33],44]"},
			List{List{Int(11)}, Int(22), List{Int(33)}, Int(44)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseJSON(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart1Solver_SolveExamplePairs(t *testing.T) {
	type args struct {
		scanner *bufio.Scanner
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"first pair: right order",
			args{
				bufio.NewScanner(strings.NewReader(
					`[1,1,3,1,1]
[1,1,5,1,1]`)),
			},
			1,
			assert.NoError,
		},
		{
			"second pair: right order",
			args{
				bufio.NewScanner(strings.NewReader(
					`[[1],[2,3,4]]
[[1],4]`)),
			},
			1,
			assert.NoError,
		},
		{
			"third pair: wrong order",
			args{
				bufio.NewScanner(strings.NewReader(
					`[9]
[[8,7,6]]`)),
			},
			0,
			assert.NoError,
		},
		{
			"fourth pair: right order",
			args{
				bufio.NewScanner(strings.NewReader(
					`[[4,4],4,4]
[[4,4],4,4,4]`,
				)),
			},
			1,
			assert.NoError,
		},
		{
			"fifth pair: wrong order",
			args{
				bufio.NewScanner(strings.NewReader(
					`[7,7,7,7]
[7,7,7]`,
				)),
			},
			0,
			assert.NoError,
		},
		{
			"sixth pair: right order",
			args{
				bufio.NewScanner(strings.NewReader(
					`[]
[3]`,
				)),
			},
			1,
			assert.NoError,
		},
		{
			"seventh pair: wrong order",
			args{
				bufio.NewScanner(strings.NewReader(
					`[[[]]]
[[]]`,
				)),
			},
			0,
			assert.NoError,
		},
		{
			"eighth pair: wrong order",
			args{
				bufio.NewScanner(strings.NewReader(
					`[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`,
				)),
			},
			0,
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Part1Solver{}
			got, err := ds.Solve(tt.args.scanner)
			if !tt.wantErr(t, err, fmt.Sprintf("Solve(%v)", tt.args.scanner)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Solve(%v)", tt.args.scanner)
		})
	}
}

func TestPart1Solver_SolveExample(t *testing.T) {
	input := `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`

	ds := &Part1Solver{}
	got, err := ds.Solve(bufio.NewScanner(strings.NewReader(input)))
	assert.NoError(t, err)
	assert.Equal(t, 13, got)
}

func TestPart1Solver_SolveMoreExamplePairs(t *testing.T) {
	type args struct {
		scanner *bufio.Scanner
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"first pair: right order",
			args{
				bufio.NewScanner(strings.NewReader(
					`[[2,[]]]
[[[]]]`)),
			},
			0,
			assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &Part1Solver{}
			got, err := ds.Solve(tt.args.scanner)
			if !tt.wantErr(t, err, fmt.Sprintf("Solve(%v)", tt.args.scanner)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Solve(%v)", tt.args.scanner)
		})
	}
}
