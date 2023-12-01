package dayXX

import (
	"bufio"
	"fmt"
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
