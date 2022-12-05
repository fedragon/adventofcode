package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/fedragon/adventofcode/day05/shared"
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
	crane := shared.Crane{
		Stacks: make(map[int]*shared.Stack),
		Moves:  make([]shared.Move, 0),
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
			parseStackLine(&crane, line)
		} else {
			move, err := parseMoveLine(line)
			if err != nil {
				panic(err)
			}

			crane.AddMove(move)
		}
	}
	crane.Finalize()

	crane.PrintStacks()
	if err := crane.ExecutePlan(&shared.CrateMover9001{}); err != nil {
		panic(err)
	}
	crane.PrintStacks()

	fmt.Printf("top crates: %v\n\n", crane.TopCrates())
}

func parseStackLine(mover *shared.Crane, line string) {
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

func parseMoveLine(line string) (shared.Move, error) {
	matches := moveRegex.FindStringSubmatch(line)
	if matches == nil {
		return shared.Move{}, fmt.Errorf("no matches in %s", line)
	}

	amount, err := strconv.Atoi(matches[moveRegex.SubexpIndex("amount")])
	if err != nil {
		return shared.Move{}, err
	}

	from, err := strconv.Atoi(matches[moveRegex.SubexpIndex("from")])
	if err != nil {
		return shared.Move{}, err
	}

	to, err := strconv.Atoi(matches[moveRegex.SubexpIndex("to")])
	if err != nil {
		return shared.Move{}, err
	}

	return shared.Move{Count: amount, From: from, To: to}, nil
}
