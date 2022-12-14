package common

import (
	"bufio"
)

type Solution struct {
	IntValue    int
	StringValue string
}

type Solver interface {
	Solve(*bufio.Scanner) (Solution, error)
}
