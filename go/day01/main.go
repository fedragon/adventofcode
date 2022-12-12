package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Calories int
}

type Elves []Elf

func (e Elves) Len() int {
	return len(e)
}

func (e Elves) Less(i, j int) bool {
	return e[i].Calories < e[j].Calories
}

func (e Elves) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

type Part1Solver struct{}

func (ds *Part1Solver) Solve(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var current Elf
	var elves Elves

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			elves = append(elves, current)
			current = Elf{}
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		current.Calories += calories
	}

	sort.Sort(sort.Reverse(elves))

	return elves[0].Calories, nil
}

type Part2Solver struct{}

func (ds *Part2Solver) Solve(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	var current Elf
	var elves Elves

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			elves = append(elves, current)
			current = Elf{}
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

	return sum, nil
}

func main() {
	part1 := Part1Solver{}
	solution, err := part1.Solve("../data/day01")
	if err != nil {
		panic(err)
	}

	fmt.Println("solution for part 1", solution)

	part2 := Part2Solver{}
	solution, err = part2.Solve("../data/day01")
	if err != nil {
		panic(err)
	}

	fmt.Println("solution for part 2", solution)
}
