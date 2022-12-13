package main

import (
	"bufio"
	"fmt"
	"github.com/fedragon/adventofcode/common"
	"os"
)

type PartXSolver struct{}

func (ds *PartXSolver) Solve(scanner *bufio.Scanner) (int, error) {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// TODO
	return 0, nil
}

func main() {
	f := common.Must(os.Open("../data/dayXX"))
	defer f.Close()

	part1 := PartXSolver{}
	solution, err := part1.Solve(bufio.NewScanner(f))
	if err != nil {
		panic(err)
	}

	fmt.Println("solution for part 1", solution)

	common.Must(f.Seek(0, 0))

	part2 := PartXSolver{}
	solution, err = part2.Solve(bufio.NewScanner(f))
	if err != nil {
		panic(err)
	}

	fmt.Println("solution for part 2", solution)
}
