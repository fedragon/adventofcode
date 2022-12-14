package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fedragon/adventofcode/common"
	"math"
	"os"
)

type Node struct {
	Coord common.Point
	Steps int
}

type Part1Solver struct{}

func (ds *Part1Solver) Solve(scanner *bufio.Scanner) (int, error) {
	var heightmap [][]rune
	var start, end common.Point

	var rowIndex int
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if len(line) == 0 {
			continue
		}

		var row []rune
		for colIndex, value := range line {
			v := value
			if value == 'S' {
				start = common.Point{X: colIndex, Y: rowIndex}
				v = 'a'
			} else if value == 'E' {
				end = common.Point{X: colIndex, Y: rowIndex}
				v = 'z'
			}

			row = append(row, v)
		}
		rowIndex++
		heightmap = append(heightmap, row)
	}

	if steps, found := visit(heightmap, Node{start, 0}, end); found {
		return steps, nil
	}

	return 0, errors.New("path not found")
}

type Part2Solver struct{}

func (ds *Part2Solver) Solve(scanner *bufio.Scanner) (int, error) {
	var heightmap [][]rune
	var end common.Point
	var starts []common.Point

	var rowIndex int
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if len(line) == 0 {
			continue
		}

		var row []rune
		for colIndex, value := range line {
			v := value
			if value == 'S' || value == 'a' {
				starts = append(starts, common.Point{X: colIndex, Y: rowIndex})
				v = 'a'
			} else if value == 'E' {
				end = common.Point{X: colIndex, Y: rowIndex}
				v = 'z'
			}

			row = append(row, v)
		}
		rowIndex++
		heightmap = append(heightmap, row)
	}

	min := math.MaxInt
	for _, start := range starts {
		if steps, found := visit(heightmap, Node{start, 0}, end); found {
			if steps < min {
				min = steps
			}
		}
	}

	if min == math.MaxInt {
		return 0, errors.New("no path found")
	}

	return min, nil
}

func alreadyQueued(xs []Node, c common.Point) bool {
	for _, x := range xs {
		if x.Coord == c {
			return true
		}
	}

	return false
}

func inGrid(c common.Point, maxRow, maxCol int) bool {
	return c.Y > -1 && c.Y < maxRow &&
		c.X > -1 && c.X < maxCol
}

func reachable(current, next rune) bool {
	return next-current <= 1
}

func visit(heightmap [][]rune, start Node, end common.Point) (int, bool) {
	q := []Node{start}
	visited := make(map[common.Point]bool)
	visited[start.Coord] = true

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		coord := curr.Coord
		visited[coord] = true

		if coord == end {
			return curr.Steps, true
		}

		value := heightmap[coord.Y][coord.X]
		neighbours := []common.Point{
			{coord.X - 1, coord.Y},
			{coord.X + 1, coord.Y},
			{coord.X, coord.Y - 1},
			{coord.X, coord.Y + 1},
		}

		for _, n := range neighbours {
			if inGrid(n, len(heightmap), len(heightmap[coord.Y])) &&
				reachable(value, heightmap[n.Y][n.X]) &&
				!visited[n] &&
				!alreadyQueued(q, n) {
				q = append(q, Node{Coord: n, Steps: curr.Steps + 1})
			}
		}
	}

	return 0, false
}

func main() {
	f := common.Must(os.Open("../data/day12"))
	defer f.Close()

	part1 := Part1Solver{}
	solution := common.Must(part1.Solve(bufio.NewScanner(f)))

	fmt.Println("solution for part 1", solution)

	common.Must(f.Seek(0, 0))

	part2 := Part2Solver{}
	solution = common.Must(part2.Solve(bufio.NewScanner(f)))

	fmt.Println("solution for part 2", solution)
}
