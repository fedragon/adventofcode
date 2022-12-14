package day04

import (
	"bufio"
	"fmt"
	"github.com/fedragon/adventofcode/common"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
}

func parse(pair string) (*Range, error) {
	tokens := strings.Split(pair, "-")
	if len(tokens) != 2 {
		return nil, fmt.Errorf("expected an assignment pair, but got %v", tokens)
	}
	start, err := strconv.Atoi(tokens[0])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(tokens[1])
	if err != nil {
		return nil, err
	}

	return &Range{Start: start, End: end}, nil
}

func (r *Range) Overlaps(a *Range) bool {
	if r.End < a.Start {
		return false
	} else if a.End < r.Start {
		return false
	}

	return true
}

func (r *Range) FullyContains(a *Range) bool {
	return r.Start <= a.Start && r.End >= a.End
}

type Part1Solver struct{}

func (s *Part1Solver) Solve(scanner *bufio.Scanner) (common.Solution, error) {
	var fullyContained int

	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, ",")
		left, err := parse(pairs[0])
		if err != nil {
			return common.Solution{}, err
		}
		right, err := parse(pairs[1])
		if err != nil {
			return common.Solution{}, err
		}

		if left.FullyContains(right) || right.FullyContains(left) {
			fullyContained++
		}
	}

	return common.Solution{IntValue: fullyContained}, nil
}

type Part2Solver struct{}

func (s *Part2Solver) Solve(scanner *bufio.Scanner) (common.Solution, error) {
	var overlapping int

	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, ",")
		left, err := parse(pairs[0])
		if err != nil {
			return common.Solution{}, err
		}
		right, err := parse(pairs[1])
		if err != nil {
			return common.Solution{}, err
		}

		if left.Overlaps(right) {
			overlapping++
		}
	}

	return common.Solution{IntValue: overlapping}, nil
}
