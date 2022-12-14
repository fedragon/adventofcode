package day01

import (
	"bufio"
	"github.com/fedragon/adventofcode/common"
	"sort"
	"strconv"
)

type elf struct {
	Calories int
}

type sortableElves []elf

func (e sortableElves) Len() int {
	return len(e)
}

func (e sortableElves) Less(i, j int) bool {
	return e[i].Calories < e[j].Calories
}

func (e sortableElves) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type Part1Solver struct{}

func (ds *Part1Solver) Solve(scanner *bufio.Scanner) (common.Solution, error) {
	var current elf
	var elves sortableElves

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			elves = append(elves, current)
			current = elf{}
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		current.Calories += calories
	}

	sort.Sort(sort.Reverse(elves))

	return common.Solution{IntValue: elves[0].Calories}, nil
}

type Part2Solver struct{}

func (ds *Part2Solver) Solve(scanner *bufio.Scanner) (common.Solution, error) {
	var current elf
	var elves sortableElves

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			elves = append(elves, current)
			current = elf{}
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		current.Calories += calories
	}

	sort.Sort(sort.Reverse(elves))

	sum := 0
	for _, e := range elves[0:3] {
		sum += e.Calories
	}

	return common.Solution{IntValue: sum}, nil
}
