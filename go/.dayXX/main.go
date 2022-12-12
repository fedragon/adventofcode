package main

import (
	"bufio"
	"fmt"
	"os"
)

type PartXSolver struct{}

func (ds *PartXSolver) Solve(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// TODO
	return 0, nil
}

func main() {
	part1 := PartXSolver{}
	solution, err := part1.Solve("../data/dayXX")
	if err != nil {
		panic(err)
	}

	fmt.Println("solution for part 1", solution)

	part2 := PartXSolver{}
	solution, err = part2.Solve("../data/dayXX")
	if err != nil {
		panic(err)
	}

	fmt.Println("solution for part 2", solution)
}
