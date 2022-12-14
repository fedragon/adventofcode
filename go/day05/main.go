package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	. "github.com/fedragon/adventofcode/day05/types"
)

func main() {
	f, err := os.Open("../data/day05")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	allStacksParsed := false
	crane9000 := Crane{
		Stacks: make(map[int]*Stack),
		Moves:  make([]Move, 0),
	}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if line[:2] == " 1" {
			allStacksParsed = true
			continue
		}

		if !allStacksParsed {
			parseStackLine(&crane9000, line)
		} else {
			move, err := parseMoveLine(line)
			if err != nil {
				panic(err)
			}

			crane9000.AddMove(move)
		}
	}
	crane9000.Finalize()

	crane9001 := crane9000.Copy()

	crane9000.PrintStacks()
	if err := crane9000.ExecutePlan(&CrateMover9000{}); err != nil {
		panic(err)
	}
	crane9000.PrintStacks()

	fmt.Printf("[CrateMover9000] Top crates: %v\n\n", crane9000.TopCrates())

	crane9001.PrintStacks()
	if err := crane9001.ExecutePlan(&CrateMover9001{}); err != nil {
		panic(err)
	}
	crane9001.PrintStacks()

	fmt.Printf("[CrateMover9001] Top crates: %v\n\n", crane9001.TopCrates())
}

func parseStackLine(mover *Crane, line string) {
	nextCrate := ""
	stackNumber := 1

	for len(line) > 0 {
		if len(line) < 4 {
			nextCrate = strings.TrimSpace(line)
			line = ""
		} else {
			nextCrate = strings.TrimSpace(line[:4])
			line = line[4:]
		}
		nextCrate = strings.Replace(strings.Replace(nextCrate, "[", "", 1), "]", "", 1)

		if len(nextCrate) > 0 {
			mover.AddCrate(stackNumber, nextCrate)
		}

		stackNumber++
	}
}

var moveRegex = regexp.MustCompile(`move (?P<amount>\d+) from (?P<from>\d+) to (?P<to>\d+)`)

func parseMoveLine(line string) (Move, error) {
	matches := moveRegex.FindStringSubmatch(line)
	if matches == nil {
		return Move{}, fmt.Errorf("no matches in %s", line)
	}

	amount, err := strconv.Atoi(matches[moveRegex.SubexpIndex("amount")])
	if err != nil {
		return Move{}, err
	}

	from, err := strconv.Atoi(matches[moveRegex.SubexpIndex("from")])
	if err != nil {
		return Move{}, err
	}

	to, err := strconv.Atoi(matches[moveRegex.SubexpIndex("to")])
	if err != nil {
		return Move{}, err
	}

	return Move{Count: amount, From: from, To: to}, nil
}
