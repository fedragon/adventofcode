package day03

import (
	"bufio"
	"github.com/fedragon/adventofcode/common"
)

type container map[rune]struct{}

func newContainer(items string) container {
	c := make(container)
	for _, i := range items {
		c[i] = struct{}{}
	}

	return c
}

func (c container) Intersect(other container) container {
	result := make(container)
	for k := range c {
		if _, ok := other[k]; ok {
			result[k] = struct{}{}
		}
	}

	return result
}

func sumBadges(lines []string) int {
	var sum int
	var group []container
	for i, line := range lines {
		if i > 0 && i%3 == 0 {
			sum += badgeOf(group)
			group = []container{}
		}

		group = append(group, newContainer(line))
	}
	sum += badgeOf(group)

	return sum
}

func badgeOf(group []container) int {
	var a container
	for _, c := range group {
		if len(a) == 0 {
			a = c
		} else {
			a = a.Intersect(c)
		}
	}

	var sum int
	for k := range a {
		sum += priority(k)
	}

	return sum
}

func priority(r rune) int {
	if r >= 97 {
		return 1 + int(r) - 97
	}

	return 27 + int(r) - 65
}

type Part1Solver struct {
}

func (s *Part1Solver) Solve(scanner *bufio.Scanner) (common.Solution, error) {
	var overlapping int

	for scanner.Scan() {
		line := scanner.Text()

		half := len(line) / 2
		first, second := newContainer(line[:half]), newContainer(line[half:])
		for item := range first.Intersect(second) {
			overlapping += priority(item)
		}
	}

	return common.Solution{IntValue: overlapping}, nil
}

type Part2Solver struct{}

func (s *Part2Solver) Solve(scanner *bufio.Scanner) (common.Solution, error) {
	var i, sum int
	var group []container

	for scanner.Scan() {
		line := scanner.Text()

		if i > 0 && i%3 == 0 {
			sum += badgeOf(group)
			group = []container{}
		}
		group = append(group, newContainer(line))
		i++
	}
	sum += badgeOf(group)

	return common.Solution{IntValue: sum}, nil
}
