package test

import (
	"bufio"
	"fmt"
	"github.com/fedragon/adventofcode/common"
	"github.com/fedragon/adventofcode/day01/day01"
	"github.com/fedragon/adventofcode/day02/day02"
	"github.com/fedragon/adventofcode/day03/day03"
	"github.com/fedragon/adventofcode/day04/day04"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func run(day int, solver common.Solver) (common.Solution, error) {
	f, err := os.Open(fmt.Sprintf("../../data/day%02d", day))
	if err != nil {
		return common.Solution{}, err
	}
	defer f.Close()

	return solver.Solve(bufio.NewScanner(f))
}

func TestDay01(t *testing.T) {
	day := 1

	solution, err := run(day, &day01.Part1Solver{})
	if assert.NoError(t, err) {
		assert.Equal(t, 69795, solution.IntValue)
	}

	solution, err = run(day, &day01.Part2Solver{})
	if assert.NoError(t, err) {
		assert.Equal(t, 208437, solution.IntValue)
	}
}

func TestDay02(t *testing.T) {
	day := 2

	solution, err := run(day, &day02.Part1Solver{})
	if assert.NoError(t, err) {
		assert.Equal(t, 15523, solution.IntValue)
	}

	solution, err = run(day, &day02.Part2Solver{})
	if assert.NoError(t, err) {
		assert.Equal(t, 15702, solution.IntValue)
	}
}

func TestDay03(t *testing.T) {
	day := 3

	solution, err := run(day, &day03.Part1Solver{})
	if assert.NoError(t, err) {
		assert.Equal(t, 7597, solution.IntValue)
	}

	solution, err = run(day, &day03.Part2Solver{})
	if assert.NoError(t, err) {
		assert.Equal(t, 2607, solution.IntValue)
	}
}

func TestDay04(t *testing.T) {
	day := 4

	solution, err := run(day, &day04.Part1Solver{})
	if assert.NoError(t, err) {
		assert.Equal(t, 657, solution.IntValue)
	}

	solution, err = run(day, &day04.Part2Solver{})
	if assert.NoError(t, err) {
		assert.Equal(t, 938, solution.IntValue)
	}
}
