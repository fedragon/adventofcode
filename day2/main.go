package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Shape int

const (
	Rock Shape = iota + 1
	Paper
	Scissors

	Loss    = 0
	Draw    = 3
	Victory = 6
)

var (
	all = map[rune]Shape{
		'A': Rock,
		'B': Paper,
		'C': Scissors,
		'X': Rock,
		'Y': Paper,
		'Z': Scissors,
	}
)

type Player struct {
	Score int
}

func (p *Player) Wins(shape Shape) {
	p.Score += int(shape) + Victory
}

func (p *Player) Draws(shape Shape) {
	p.Score += int(shape) + Draw
}

func (p *Player) Loses(shape Shape) {
	p.Score += int(shape) + Loss
}

type Board struct {
	Opponent Player
	Player   Player
}

func (b *Board) NextRound(input string) {
	l, ok := all[rune(input[0])]
	if !ok {
		log.Fatalf("unknown symbol: %c", input[0])
	}
	r, ok := all[rune(input[2])]
	if !ok {
		log.Fatalf("unknown symbol: %c", input[2])
	}

	switch {
	case l == Rock && r == Scissors:
		b.Opponent.Wins(l)
		b.Player.Loses(r)
	case l == Scissors && r == Rock:
		b.Opponent.Loses(l)
		b.Player.Wins(r)
	case l > r:
		b.Opponent.Wins(l)
		b.Player.Loses(r)
	case l < r:
		b.Opponent.Loses(l)
		b.Player.Wins(r)
	default:
		b.Opponent.Draws(l)
		b.Player.Draws(r)
	}
}

func main() {
	var board Board

	f, err := os.Open("day2/input")
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
