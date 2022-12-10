package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var addxRe = regexp.MustCompile(`^addx (-?\d+)$`)

type Register struct {
	Value int
}

type Instruction struct {
	Name   string
	Cycles int
	Effect func(*Register)
}

func main() {
	f, err := os.Open("../data/day10")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var instructions []Instruction

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "noop" {
			instructions = append(instructions, Instruction{
				Name:   "noop",
				Cycles: 1,
				Effect: func(*Register) {},
			})
		} else {
			matches := addxRe.FindStringSubmatch(line)
			value, _ := strconv.Atoi(matches[1])
			instructions = append(instructions, Instruction{
				Name:   matches[0],
				Cycles: 2,
				Effect: func(r *Register) {
					r.Value += value
				},
			})
		}
	}

	var sum int
	for k, v := range Run(instructions) {
		switch k {
		case 20, 60, 100, 140, 180, 220:
			fmt.Printf("%d: %d\n", k, v)
			sum += v
		}
	}
	fmt.Println("sum", sum)
}

type CRT struct {
	pixels   [6][40]rune
	row, col int
}

func (c *CRT) Draw(x Register) {
	char := '.'
	if c.col >= x.Value-1 && c.col <= x.Value+1 {
		char = '#'
	}

	if c.row > 6 {
		return
	}
	if c.col > 39 {
		c.row++
		c.col = 0
	}

	c.pixels[c.row][c.col] = char
	c.col++
}

func (c *CRT) Print() {
	for _, row := range c.pixels {
		for _, p := range row {
			if p == 0 {
				fmt.Printf("%c", '.')
			} else {
				fmt.Printf("%c", p)
			}
		}
		fmt.Println()
	}
}

func Run(instructions []Instruction) map[int]int {
	x := Register{Value: 1}

	var cycle int
	signalStrengths := make(map[int]int)

	crt := &CRT{}
	for _, instr := range instructions {
		for c := 0; c < instr.Cycles; c++ {
			fmt.Println("instruction", instr, "register", x.Value)
			crt.Draw(x)
			cycle++

			signalStrengths[cycle] = x.Value * cycle
		}
		instr.Effect(&x)
	}

	crt.Print()

	return signalStrengths
}
