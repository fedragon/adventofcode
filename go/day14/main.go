package main

import (
	"bufio"
	"fmt"
	"github.com/fedragon/adventofcode/common"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Tile int

const (
	Air Tile = iota
	Rock
	Sand
	SandSource
)

type Grid struct {
	Tiles                  map[int]map[int]Tile
	MinX, MaxX, MinY, MaxY int
}

func (g *Grid) String() string {
	var b strings.Builder

	for r := g.MinX; r <= g.MaxX; r++ {
		for c := g.MinY; c <= g.MaxY; c++ {
			switch g.Tiles[r][c] {
			case Rock:
				b.WriteRune('#')
			case Sand:
				b.WriteRune('o')
			case SandSource:
				b.WriteRune('+')
			default:
				b.WriteRune('.')
			}
		}
		b.WriteRune('\n')
	}

	return b.String()
}

type Segment [2]common.Coord

func (s Segment) Walk() []common.Coord {
	var from, to common.Coord
	var coords []common.Coord

	if s[0].X < s[1].X || s[0].Y < s[1].Y {
		from = s[0]
		to = s[1]
	} else {
		from = s[1]
		to = s[0]
	}

	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			coords = append(coords, common.Coord{X: x, Y: y})
		}
	}

	return coords
}

type Part1Solver struct{}

func (ds *Part1Solver) Solve(scanner *bufio.Scanner) (int, error) {
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		lines = append(lines, line)
	}

	fmt.Println(buildGrid(lines))

	// TODO
	return 0, nil
}

var re = regexp.MustCompile(`(\d+,\d+)[\s\->]?`)

func buildGrid(lines []string) *Grid {
	tiles := map[int]map[int]Tile{
		0: {500: SandSource},
	}
	minX, maxX, minY, maxY := 0, 0, 500, 0

	for _, line := range lines {
		for _, segment := range parse(line) {
			for _, coord := range segment.Walk() {
				if tiles[coord.X] == nil {
					tiles[coord.X] = make(map[int]Tile)
				}

				tiles[coord.X][coord.Y] = Rock

				if coord.X < minX {
					minX = coord.X
				} else if coord.X > maxX {
					maxX = coord.X
				}

				if coord.Y < minY {
					minY = coord.Y
				} else if coord.Y > maxY {
					maxY = coord.Y
				}
			}
		}
	}

	return &Grid{Tiles: tiles, MinX: minX, MaxX: maxX, MinY: minY, MaxY: maxY}
}

func parse(line string) []Segment {
	var path []Segment
	var previous common.Coord

	for _, matches := range re.FindAllStringSubmatch(line, -1) {
		xy := strings.Split(matches[1], ",")

		if previous == (common.Coord{}) {
			previous = common.Coord{
				Y: common.Must(strconv.Atoi(xy[0])),
				X: common.Must(strconv.Atoi(xy[1])),
			}
			continue
		}

		current := common.Coord{
			Y: common.Must(strconv.Atoi(xy[0])),
			X: common.Must(strconv.Atoi(xy[1])),
		}
		path = append(path, [2]common.Coord{
			previous,
			current,
		})
		previous = current
	}

	return path
}

func main() {
	f := common.Must(os.Open("../data/day14"))
	defer f.Close()

	part1 := Part1Solver{}
	solution := common.Must(part1.Solve(bufio.NewScanner(f)))

	fmt.Println("solution for part 1", solution)

	//common.Must(f.Seek(0, 0))
	//
	//part2 := PartXSolver{}
	//solution = common.Must(part2.Solve(bufio.NewScanner(f)))
	//
	//fmt.Println("solution for part 2", solution)
}
