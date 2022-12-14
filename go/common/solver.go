package common

import "bufio"

type Solver interface {
	Solve(*bufio.Scanner) (int, error)
}
