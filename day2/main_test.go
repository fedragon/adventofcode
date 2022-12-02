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
			"left=Rock right=Loses",
			Board{},
			"A X",
			Board{Opponent: Player{Score: 7}, Player: Player{Score: 3}},
		},
		{
			"left=Paper right=Loses",
			Board{},
			"B X",
			Board{Opponent: Player{Score: 8}, Player: Player{Score: 1}},
		},
		{
			"left=Scissors right=Loses",
			Board{},
			"C X",
			Board{Opponent: Player{Score: 9}, Player: Player{Score: 2}},
		},
		{
			"left=Rock right=Draws",
			Board{},
			"A Y",
			Board{Opponent: Player{Score: 4}, Player: Player{Score: 4}},
		},
		{
			"left=Paper right=Draws",
			Board{},
			"B Y",
			Board{Opponent: Player{Score: 5}, Player: Player{Score: 5}},
		},
		{
			"left=Scissors right=Draws",
			Board{},
			"C Y",
			Board{Opponent: Player{Score: 6}, Player: Player{Score: 6}},
		},
		{
			"left=Rock right=Wins",
			Board{},
			"A Z",
			Board{Opponent: Player{Score: 1}, Player: Player{Score: 8}},
		},
		{
			"left=Paper right=Wins",
			Board{},
			"B Z",
			Board{Opponent: Player{Score: 2}, Player: Player{Score: 9}},
		},
		{
			"left=Scissors right=Wins",
			Board{},
			"C Z",
			Board{Opponent: Player{Score: 3}, Player: Player{Score: 7}},
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
