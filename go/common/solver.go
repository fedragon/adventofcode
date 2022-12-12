package common

type Solver interface {
	Solve(filename string) (int, error)
}
