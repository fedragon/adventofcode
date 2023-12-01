package main

import (
	"bufio"
	"fmt"
	"github.com/fedragon/adventofcode/.dayXX/dayXX"
	"github.com/fedragon/adventofcode/common"
	"os"
)

func main() {
	f := common.Must(os.Open("../data/dayXX"))
	defer f.Close()

	part1 := dayXX.PartXSolver{}
	solution := common.Must(part1.Solve(bufio.NewScanner(f)))

	fmt.Println("solution for part 1", solution)

	common.Must(f.Seek(0, 0))

	part2 := dayXX.PartXSolver{}
	solution = common.Must(part2.Solve(bufio.NewScanner(f)))

	fmt.Println("solution for part 2", solution)
}
