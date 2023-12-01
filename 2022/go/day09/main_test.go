package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestVisit(t *testing.T) {
	tests := []struct {
		name  string
		moves []Move
		want  History
	}{
		{
			"the tail initially visits the starting position",
			[]Move{{Direction: "U", Steps: 1}},
			History{
				0: {0: 'T'},
			},
		},
		{
			"it stays put if the head remains adjacent",
			[]Move{{Direction: "U", Steps: 1}, {Direction: "R", Steps: 1}},
			History{
				0: {0: 'T'},
			},
		},
		{
			"it follows the head up, staying one step behind whenever it is not adjacent",
			[]Move{{Direction: "U", Steps: 2}},
			History{
				-1: {0: 'T'},
				0:  {0: 'T'},
			},
		},
		{
			"it follows the head down, staying one step behind whenever it is not adjacent",
			[]Move{{Direction: "D", Steps: 2}},
			History{
				0: {0: 'T'},
				1: {0: 'T'},
			},
		},
		{
			"it follows the head left, staying one step behind whenever it is not adjacent",
			[]Move{{Direction: "L", Steps: 2}},
			History{
				0: {-1: 'T', 0: 'T'},
			},
		},
		{
			"it follows the head right, staying one step behind whenever it is not adjacent",
			[]Move{{Direction: "R", Steps: 2}},
			History{
				0: {0: 'T', 1: 'T'},
			},
		},
		{
			"it follows the head right, staying one step behind whenever it is not adjacent",
			[]Move{{Direction: "R", Steps: 2}},
			History{
				0: {0: 'T', 1: 'T'},
			},
		},
		{
			"it stays put when the head moves on top of it",
			[]Move{{Direction: "U", Steps: 2}, {Direction: "D", Steps: 1}},
			History{
				-1: {0: 'T'},
				0:  {0: 'T'},
			},
		},
		{
			"it moves diagonally",
			[]Move{{Direction: "U", Steps: 1}, {Direction: "R", Steps: 1}, {Direction: "U", Steps: 1}},
			History{
				-1: {1: 'T'},
				0:  {0: 'T'},
			},
		},
		{
			"it solves the example in part1",
			[]Move{
				{"R", 4},
				{"U", 4},
				{"L", 3},
				{"D", 1},
				{"R", 4},
				{"D", 1},
				{"L", 5},
				{"R", 2},
			},
			History{
				-4: {2: 84, 3: 84},
				-3: {3: 84, 4: 84},
				-2: {1: 84, 2: 84, 3: 84, 4: 84},
				-1: {4: 84},
				0:  {0: 84, 1: 84, 2: 84, 3: 84}},
		},
	}

	_ = os.Setenv("LOG_TRAIL", "1")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Visit(tt.moves); !assert.EqualValues(t, tt.want, got) {
				t.Errorf("Visit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVisitN(t *testing.T) {
	type args struct {
		moves    []Move
		numKnots int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"it moves a list of knots 4 steps to the right",
			args{
				[]Move{{"R", 4}},
				4,
			},
			2,
		},
		{
			"it moves a list of knots 10 steps to the right",
			args{
				[]Move{{"R", 10}},
				10,
			},
			2,
		},
		{
			"it moves a list of knots 3 steps to the right and then up",
			args{
				[]Move{{"R", 3}, {"U", 3}},
				3,
			},
			3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			knots := make([]*Knot, tt.args.numKnots)
			for i := 0; i < tt.args.numKnots; i++ {
				knots[i] = &Knot{rune('0' + i), 0, 0, map[string]struct{}{"(0,0)": {}}}
			}
			VisitN(knots, tt.args.moves)

			if got := len(knots[tt.args.numKnots-1].Visited); !assert.EqualValues(t, tt.want, got) {
				t.Errorf("VisitN() = %v, want %v", got, tt.want)
			}
		})
	}
}
