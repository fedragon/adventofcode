package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Calories int
}

type Elves []Elf

func (e Elves) Len() int {
	return len(e)
}

func (e Elves) Less(i, j int) bool {
	return e[i].Calories < e[j].Calories
}

func (e Elves) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func main() {
	var current Elf
	var elves Elves

	f, err := os.Open("../data/day1")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			elves = append(elves, current)
			current = Elf{}
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		current.Calories += calories
	}

	sort.Sort(sort.Reverse(elves))

	sum := 0
	for _, e := range elves[0:3] {
		sum += e.Calories
	}

	fmt.Println(sum)
}
