package day02

import (
	"bufio"
	"fmt"
	"github.com/fedragon/adventofcode/common"
)

type shape int
type result rune

const (
	Rock shape = iota + 1
	Paper
	Scissors

	Loss    result = 'X'
	Draw    result = 'Y'
	Victory result = 'Z'
)

var (
	shapes = map[rune]shape{
		'A': Rock,
		'B': Paper,
		'C': Scissors,
		'X': Rock,
		'Y': Paper,
		'Z': Scissors,
	}
)

type player struct {
	Score int
}

func (p *player) Wins(shape shape) {
	p.Score += int(shape) + 6
}

func (p *player) Draws(shape shape) {
	p.Score += int(shape) + 3
}

func (p *player) Loses(shape shape) {
	p.Score += int(shape) + 0
}

type Part1Solver struct {
	Opponent player
	Player   player
}

func (s *Part1Solver) Solve(scanner *bufio.Scanner) (common.Solution, error) {
	nextRound := func(input string) error {
		l, ok := shapes[rune(input[0])]
		if !ok {
			return fmt.Errorf("unknown symbol: %c", input[0])
		}
		r, ok := shapes[rune(input[2])]
		if !ok {
			return fmt.Errorf("unknown symbol: %c", input[2])
		}

		switch {
		case l == Rock && r == Scissors:
			s.Opponent.Wins(l)
			s.Player.Loses(r)
		case l == Scissors && r == Rock:
			s.Opponent.Loses(l)
			s.Player.Wins(r)
		case l > r:
			s.Opponent.Wins(l)
			s.Player.Loses(r)
		case l < r:
			s.Opponent.Loses(l)
			s.Player.Wins(r)
		default:
			s.Opponent.Draws(l)
			s.Player.Draws(r)
		}

		return nil
	}

	for scanner.Scan() {
		if err := nextRound(scanner.Text()); err != nil {
			return common.Solution{}, err
		}
	}

	return common.Solution{IntValue: s.Player.Score}, nil
}

type Part2Solver struct {
	Opponent player
	Player   player
}

func (s *Part2Solver) Solve(scanner *bufio.Scanner) (common.Solution, error) {
	nextRound := func(input string) error {
		opponentMove, ok := shapes[rune(input[0])]
		if !ok {
			return fmt.Errorf("unknown symbol: %c", input[0])
		}
		expectedResult := result(input[2])

		switch {
		case expectedResult == Draw:
			s.Opponent.Draws(opponentMove)
			s.Player.Draws(opponentMove)
		case expectedResult == Victory:
			playerMove := opponentMove + 1
			if playerMove > Scissors {
				playerMove = Rock
			}
			s.Opponent.Loses(opponentMove)
			s.Player.Wins(playerMove)
		case expectedResult == Loss:
			playerMove := opponentMove - 1
			if playerMove < Rock {
				playerMove = Scissors
			}
			s.Opponent.Wins(opponentMove)
			s.Player.Loses(playerMove)
		}

		return nil
	}

	for scanner.Scan() {
		if err := nextRound(scanner.Text()); err != nil {
			return common.Solution{}, err
		}
	}

	return common.Solution{IntValue: s.Player.Score}, nil
}
