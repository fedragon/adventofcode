package day15

import (
	"bufio"
	"fmt"
	"github.com/fedragon/adventofcode/common"
	"regexp"
	"strconv"
	"strings"
)

const (
	Beacon   = 'B'
	Sensor   = 'S'
	NoSensor = '#'
)

// Grid
//
//	TODO this is VERY similar to the grid in day14, can I create a more generic one and reuse it?
type Grid struct {
	Tiles      map[common.Point]rune
	MinX, MaxX int
	MinY, MaxY int
}

func (g *Grid) String() string {
	var b strings.Builder
	for r := g.MinY; r <= g.MaxY; r++ {
		for c := g.MinX; c <= g.MaxX; c++ {
			r, ok := g.Tiles[common.Point{X: c, Y: r}]
			if ok {
				b.WriteRune(r)
			} else {
				b.WriteRune('.')
			}
		}
		b.WriteRune('\n')
	}

	return b.String()
}

var re = regexp.MustCompile(`x=(-?\d+), y=(-?\d+)`)

func minmax(min, max, next int) (int, int) {
	if next < min {
		min = next
	} else if next > max {
		max = next
	}

	return min, max
}

func buildGrid(lines []string) *Grid {
	tiles := map[common.Point]rune{}
	minY, maxY := 0, 0
	minX, maxX := 0, 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		matches := re.FindAllStringSubmatch(line, -1)
		fmt.Println(matches)

		sensor := common.Point{
			X: common.Must(strconv.Atoi(matches[0][1])),
			Y: common.Must(strconv.Atoi(matches[0][2])),
		}
		tiles[sensor] = Sensor

		minX, maxX = minmax(minX, maxX, sensor.X)
		minY, maxY = minmax(minY, maxY, sensor.Y)

		beacon := common.Point{
			X: common.Must(strconv.Atoi(matches[1][1])),
			Y: common.Must(strconv.Atoi(matches[1][2])),
		}
		tiles[beacon] = Beacon

		minX, maxX = minmax(minX, maxX, beacon.X)
		minY, maxY = minmax(minY, maxY, beacon.Y)
	}

	return &Grid{Tiles: tiles, MinX: minX, MaxX: maxX, MinY: minY, MaxY: maxY}
}

type Part1Solver struct{}

func (ds *Part1Solver) Solve(scanner *bufio.Scanner) (int, error) {
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	grid := buildGrid(lines)
	fmt.Println(grid.String())

	return 0, nil
}

type Part2Solver struct{}

func (ds *Part2Solver) Solve(scanner *bufio.Scanner) (int, error) {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// TODO
	return 0, nil
}
