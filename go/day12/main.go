package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
)

var emptyNode = Node{}

type Coord struct {
	Row, Col int
}

type Node struct {
	Value    rune
	Coord    Coord
	Distance int
}

func (n *Node) Visit(grid [][]Node) (int, bool) {
	return visit(grid, *n, nil)
}

// neighboursIn returns a list of nodes sorted by elevationGain in descending order
func (n *Node) neighboursIn(grid [][]Node) []Node {
	var neighbours []Node

	exists := func(c Coord) bool {
		return c.Row > -1 && c.Row < len(grid) &&
			c.Col > -1 && c.Col < len(grid[c.Row])
	}

	reachable := func(current, next rune) bool {
		return (current == 'S' && next == 'a') ||
			(current == 'z' && next == 'E') ||
			next-current <= 1
	}

	move := func(row, col int) []Coord {
		return []Coord{
			{row - 1, col},
			{row + 1, col},
			{row, col - 1},
			{row, col + 1},
		}
	}

	gains := make(map[int][]Node)
	for _, coord := range move(n.Coord.Row, n.Coord.Col) {
		if exists(coord) {
			if next := grid[coord.Row][coord.Col]; reachable(n.Value, next.Value) {
				neighbours = append(neighbours, next)
				elevationGain := int(next.Value - n.Value)
				gains[elevationGain] = append(gains[elevationGain], next)
			}
		}
	}

	var keys []int
	for k := range gains {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var sorted []Node
	for _, k := range keys {
		sorted = append(sorted, gains[k]...)
	}

	return sorted
}

func inVisited(visited []Node, n Node) bool {
	for _, x := range visited {
		if n.Value == x.Value && n.Coord.Row == x.Coord.Row && n.Coord.Col == x.Coord.Col {
			return true
		}
	}

	return false
}

func visit(grid [][]Node, start Node, visited []Node) (int, bool) {
	if start == emptyNode {
		return math.MaxInt, false
	}

	neighbours := start.neighboursIn(grid)
	fmt.Printf("start: %v (%d, %d): %d, len(visited) = %d\n", start.Value, start.Coord.Row, start.Coord.Col, start.Distance, len(visited))

	for i := range neighbours {
		if neighbours[i].Value == 'E' {
			neighbours[i].Distance = start.Distance + 1
			return start.Distance + 1, true
		}

		if !inVisited(visited, neighbours[i]) {
			neighbours[i].Distance = start.Distance + 1
		}
	}

	for i := range neighbours {
		if !inVisited(visited, neighbours[i]) {
			if res, found := visit(grid, neighbours[i], append(visited, start)); found {
				return res, true
			}

			visited = append(visited, neighbours[i])
		}
	}

	return math.MaxInt, false
}

type Part1Solver struct{}

func (ds *Part1Solver) Solve(scanner *bufio.Scanner) (int, error) {
	var heightmap [][]Node
	var start Node

	var rowIndex int
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if len(line) == 0 {
			continue
		}

		var row []Node
		for colIndex, value := range line {
			node := Node{
				Value: value,
				Coord: Coord{
					Row: rowIndex,
					Col: colIndex,
				},
				Distance: math.MaxInt,
			}

			if value == 'S' {
				node.Distance = 0
				start = node
			}

			row = append(row, node)
		}
		rowIndex++
		heightmap = append(heightmap, row)
	}

	if steps, found := start.Visit(heightmap); found {
		return steps, nil
	}

	return 0, errors.New("no path found")
}

func main() {
	f, err := os.Open("../data/day12")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	part1 := Part1Solver{}
	solution, err := part1.Solve(scanner)
	if err != nil {
		panic(err)
	}

	fmt.Println("solution for part 1", solution)

	// part2 := PartXSolver{}
	// solution, err = part2.Solve("../data/day12")
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//
	// fmt.Println("solution for part 2", solution)
}
