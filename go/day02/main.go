package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Shape int
type Result rune

const (
	Rock Shape = iota + 1
	Paper
	Scissors

	Loss    Result = 'X'
	Draw    Result = 'Y'
	Victory Result = 'Z'
)

var (
	shapes = map[rune]Shape{
		'A': Rock,
		'B': Paper,
		'C': Scissors,
	}
)

type Player struct {
	Score int
}

func (p *Player) Wins(shape Shape) {
	p.Score += int(shape) + 6
}

func (p *Player) Draws(shape Shape) {
	p.Score += int(shape) + 3
}

func (p *Player) Loses(shape Shape) {
	p.Score += int(shape) + 0
}

type Board struct {
	Opponent Player
	Player   Player
}

func (b *Board) NextRound(input string) {
	opponentMove, ok := shapes[rune(input[0])]
	if !ok {
		log.Fatalf("unknown symbol: %c", input[0])
	}
	expectedResult := Result(input[2])

	switch {
	case expectedResult == Draw:
		b.Opponent.Draws(opponentMove)
		b.Player.Draws(opponentMove)
	case expectedResult == Victory:
		playerMove := opponentMove + 1
		if playerMove > Scissors {
			playerMove = Rock
		}
		b.Opponent.Loses(opponentMove)
		b.Player.Wins(playerMove)
	case expectedResult == Loss:
		playerMove := opponentMove - 1
		if playerMove < Rock {
			playerMove = Scissors
		}
		b.Opponent.Wins(opponentMove)
		b.Player.Loses(playerMove)
	}
}

func main() {
	var board Board

	f, err := os.Open("../data/day02")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		board.NextRound(scanner.Text())
	}

	fmt.Printf("%+v\n", board)
}
