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
	Source

	SourceX = 500
	SourceY = 0
)

type Grid struct {
	Tiles      map[common.Point]Tile
	MinX, MaxX int
	MinY, MaxY int
	Floor      *int
}

func (g *Grid) IsEmpty(p common.Point) bool {
	if g.Floor == nil {
		return g.Tiles[p] == Air
	}

	if p.X == SourceX && p.Y == SourceY && g.Tiles[p] == Source {
		return true
	}

	if p.Y < *g.Floor {
		return g.Tiles[p] == Air
	}

	return false
}

func (g *Grid) TileAt(p common.Point) Tile {
	if g.Floor == nil || p.Y < *g.Floor {
		return g.Tiles[p]
	}

	return Rock
}

func (g *Grid) PourSand(x, y int) bool {
	max := g.MaxY
	if g.Floor != nil {
		max = *g.Floor
	}

	for y <= max && g.IsEmpty(common.Point{X: x, Y: y}) {
		if x < g.MinX {
			g.MinX = x - 1
		} else if x > g.MaxX {
			g.MaxX = x + 1
		}

		if g.IsEmpty(common.Point{Y: y + 1, X: x}) {
			y++
		} else if g.IsEmpty(common.Point{Y: y + 1, X: x - 1}) {
			y++
			x--
		} else if g.IsEmpty(common.Point{Y: y + 1, X: x + 1}) {
			y++
			x++
		} else {
			if y == 0 {
				fmt.Println()
			}

			g.Tiles[common.Point{X: x, Y: y}] = Sand
			return true
		}
	}

	return false
}

func (g *Grid) String() string {
	max := g.MaxY
	if g.Floor != nil {
		max = *g.Floor
	}

	var b strings.Builder
	for r := g.MinY; r <= max; r++ {
		for c := g.MinX; c <= g.MaxX; c++ {
			switch g.TileAt(common.Point{X: c, Y: r}) {
			case Rock:
				b.WriteRune('#')
			case Sand:
				b.WriteRune('o')
			case Source:
				b.WriteRune('+')
			default:
				b.WriteRune('.')
			}
		}
		b.WriteRune('\n')
	}

	return b.String()
}

type Segment [2]common.Point

func (s Segment) Walk() []common.Point {
	var from, to common.Point
	var coords []common.Point

	if s[0].X < s[1].X || s[0].Y < s[1].Y {
		from = s[0]
		to = s[1]
	} else {
		from = s[1]
		to = s[0]
	}

	for x := from.X; x <= to.X; x++ {
		for y := from.Y; y <= to.Y; y++ {
			coords = append(coords, common.Point{X: x, Y: y})
		}
	}

	return coords
}

type Part1Solver struct{}

func (ds *Part1Solver) Solve(scanner *bufio.Scanner) (int, error) {
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	grid := buildGrid(lines)
	var count int

	for {
		if ok := grid.PourSand(SourceX, SourceY+1); !ok {
			break
		}
		count++
	}

	return count, nil
}

type Part2Solver struct{}

func (ds *Part2Solver) Solve(scanner *bufio.Scanner) (int, error) {
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	trail := common.Must(os.Create("trail.txt"))
	defer trail.Close()

	grid := buildGrid(lines)
	floor := grid.MaxY + 2
	grid.Floor = &floor

	var count int

	trail.WriteString(grid.String())
	trail.WriteString("\n====================\n")

	for {
		if ok := grid.PourSand(SourceX, SourceY); !ok {
			break
		}
		count++
	}

	trail.WriteString(grid.String())

	return count, nil
}

var re = regexp.MustCompile(`(\d+,\d+)[\s\->]?`)

func buildGrid(lines []string) *Grid {
	tiles := map[common.Point]Tile{
		common.Point{X: SourceX, Y: SourceY}: Source,
	}
	minY, maxY := 0, SourceY
	minX, maxX := SourceX, 0

	for _, line := range lines {
		for _, segment := range parse(line) {
			for _, coord := range segment.Walk() {
				tiles[coord] = Rock

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
	var previous common.Point

	for _, matches := range re.FindAllStringSubmatch(line, -1) {
		xy := strings.Split(matches[1], ",")

		if previous == (common.Point{}) {
			previous = common.Point{
				X: common.Must(strconv.Atoi(xy[0])),
				Y: common.Must(strconv.Atoi(xy[1])),
			}
			continue
		}

		current := common.Point{
			X: common.Must(strconv.Atoi(xy[0])),
			Y: common.Must(strconv.Atoi(xy[1])),
		}
		path = append(path, [2]common.Point{
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

	common.Must(f.Seek(0, 0))

	part2 := Part2Solver{}
	solution = common.Must(part2.Solve(bufio.NewScanner(f)))

	fmt.Println("solution for part 2", solution)
}
