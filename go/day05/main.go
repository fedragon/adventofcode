package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type stack struct {
	elems []string
}

func (s *stack) Prepend(elem string) {
	s.elems = append([]string{elem}, s.elems...)
}

func (s *stack) Push(elem string) {
	s.elems = append(s.elems, elem)
}

func (s *stack) Pop() string {
	if len(s.elems) == 0 {
		return ""
	}

	elem := s.elems[len(s.elems)-1]
	s.elems = s.elems[:len(s.elems)-1]

	return elem
}

func (s *stack) String() string {
	return fmt.Sprint(s.elems)
}

func main() {
	f, err := os.Open("../data/day05")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	stacks := make(map[int]*stack)

	var nextCrate string
	for scanner.Scan() {
		line := scanner.Text()

		if line[:2] == " 1" {
			break
		}

		stackNumber := 1
		for len(line) > 0 {
			if len(line) < 4 {
				nextCrate = strings.TrimSpace(line)
				line = ""
			} else {
				nextCrate = strings.TrimSpace(line[:4])
				line = line[4:]
			}

			current, ok := stacks[stackNumber]
			if !ok {
				current = &stack{}
				stacks[stackNumber] = current
			}

			if len(nextCrate) > 0 {
				current.Prepend(nextCrate)
			}

			stackNumber++
		}
	}

	fmt.Println(stacks)
}
