package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fedragon/adventofcode/common"
	"math"
	"os"
)

type Coord struct {
	Row, Col int
}

type Node struct {
	Coord Coord
	Steps int
}

type Part1Solver struct{}

func (ds *Part1Solver) Solve(scanner *bufio.Scanner) (int, error) {
	var heightmap [][]rune
	var start, end Coord

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
				start = Coord{Row: rowIndex, Col: colIndex}
				v = 'a'
			} else if value == 'E' {
				end = Coord{Row: rowIndex, Col: colIndex}
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
	var end Coord
	var starts []Coord

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
				starts = append(starts, Coord{Row: rowIndex, Col: colIndex})
				v = 'a'
			} else if value == 'E' {
				end = Coord{Row: rowIndex, Col: colIndex}
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

func alreadyQueued(xs []Node, c Coord) bool {
	for _, x := range xs {
		if x.Coord == c {
			return true
		}
	}

	return false
}

func inGrid(c Coord, maxRow, maxCol int) bool {
	return c.Row > -1 && c.Row < maxRow &&
		c.Col > -1 && c.Col < maxCol
}

func reachable(current, next rune) bool {
	return next-current <= 1
}

func visit(heightmap [][]rune, start Node, end Coord) (int, bool) {
	q := []Node{start}
	visited := make(map[Coord]bool)
	visited[start.Coord] = true

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		coord := curr.Coord
		visited[coord] = true

		if coord == end {
			return curr.Steps, true
		}

		value := heightmap[coord.Row][coord.Col]
		neighbours := []Coord{
			{coord.Row - 1, coord.Col},
			{coord.Row + 1, coord.Col},
			{coord.Row, coord.Col - 1},
			{coord.Row, coord.Col + 1},
		}

		for _, n := range neighbours {
			if inGrid(n, len(heightmap), len(heightmap[coord.Row])) &&
				reachable(value, heightmap[n.Row][n.Col]) &&
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
	solution, err := part1.Solve(bufio.NewScanner(f))
	if err != nil {
		panic(err)
	}

	fmt.Println("solution for part 1", solution)

	common.Must(f.Seek(0, 0))

	part2 := Part2Solver{}
	solution, err = part2.Solve(bufio.NewScanner(f))

	if err != nil {
		panic(err)
	}

	fmt.Println("solution for part 2", solution)
}
