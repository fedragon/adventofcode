package main

import (
	"bufio"
	"fmt"
	"os"
)

type Container map[rune]struct{}

func NewContainer(items string) Container {
	c := make(Container)
	for _, i := range items {
		c[i] = struct{}{}
	}

	return c
}

func (c Container) Intersect(other Container) Container {
	result := make(Container)
	for k := range c {
		if _, ok := other[k]; ok {
			result[k] = struct{}{}
		}
	}

	return result
}

func main() {
	f, err := os.Open("../data/day03")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(sumOverlapping(lines))

	fmt.Println(sumBadges(lines))
}

func sumOverlapping(lines []string) int {
	var sum int
	for _, line := range lines {
		half := len(line) / 2
		first, second := NewContainer(line[:half]), NewContainer(line[half:])
		for item := range first.Intersect(second) {
			sum += priority(item)
		}
	}
	return sum
}

func sumBadges(lines []string) int {
	var sum int
	var group []Container
	for i, line := range lines {
		if i > 0 && i%3 == 0 {
			sum += badgeOf(group)
			group = []Container{}
		}

		group = append(group, NewContainer(line))
	}
	sum += badgeOf(group)

	return sum
}

func badgeOf(group []Container) int {
	var a Container
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
