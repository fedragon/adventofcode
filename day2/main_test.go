package main

import "testing"

func TestBoard_NextRound(t *testing.T) {
	tests := []struct {
		name     string
		board    Board
		input    string
		expected Board
	}{
		{
			"left=Scissors right=Rock",
			Board{},
			"C X",
			Board{Opponent: Player{Score: 3}, Player: Player{Score: 7}},
		},
		{
			"left=Rock right=Scissors",
			Board{},
			"A Z",
			Board{Opponent: Player{Score: 7}, Player: Player{Score: 3}},
		},
		{
			"left=Rock right=Paper",
			Board{},
			"A Y",
			Board{Opponent: Player{Score: 1}, Player: Player{Score: 8}},
		},
		{
			"left=Paper right=Rock",
			Board{},
			"B X",
			Board{Opponent: Player{Score: 8}, Player: Player{Score: 1}},
		},
		{
			"left=Paper right=Scissors",
			Board{},
			"B Z",
			Board{Opponent: Player{Score: 2}, Player: Player{Score: 9}},
		},
		{
			"left=Scissors right=Paper",
			Board{},
			"C Y",
			Board{Opponent: Player{Score: 9}, Player: Player{Score: 2}},
		},
		{
			"left=Rock right=Rock",
			Board{},
			"A X",
			Board{Opponent: Player{Score: 4}, Player: Player{Score: 4}},
		},
		{
			"left=Paper right=Paper",
			Board{},
			"B Y",
			Board{Opponent: Player{Score: 5}, Player: Player{Score: 5}},
		},
		{
			"left=Scissors right=Scissors",
			Board{},
			"C Z",
			Board{Opponent: Player{Score: 6}, Player: Player{Score: 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.board.NextRound(tt.input)

			if tt.expected != tt.board {
				t.Errorf("expected %v, but got %v", tt.expected, tt.board)
			}
		})
	}
}
