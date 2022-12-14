package main

import (
	"bufio"
	"fmt"
	"github.com/fedragon/adventofcode/common"
	"github.com/fedragon/adventofcode/day02/day02"
	"os"
)

func main() {
	f := common.Must(os.Open("../data/day02"))
	defer f.Close()

	part1 := day02.Part1Solver{}
	solution := common.Must(part1.Solve(bufio.NewScanner(f)))

	fmt.Println("solution for part 1", solution)

	common.Must(f.Seek(0, 0))

	part2 := day02.Part2Solver{}
	solution = common.Must(part2.Solve(bufio.NewScanner(f)))

	fmt.Println("solution for part 2", solution)
}
