package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
}

func Parse(pair string) (*Range, error) {
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

func main() {
	f, err := os.Open("../data/day04")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var fullyContained int
	var overlapping int
	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, ",")
		left, err := Parse(pairs[0])
		if err != nil {
			panic(err)
		}
		right, err := Parse(pairs[1])
		if err != nil {
			panic(err)
		}

		if left.FullyContains(right) || right.FullyContains(left) {
			fullyContained++
		}

		if left.Overlaps(right) {
			overlapping++
		}
	}

	fmt.Println("fully contained ranges", fullyContained)
	fmt.Println("overlapping ranges", overlapping)
}
