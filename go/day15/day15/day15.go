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

type Grid struct {
	Tiles      map[common.Point]rune
	MinX, MaxX int
	MinY, MaxY int
}

func (g *Grid) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("min(y,x) = (%d, %d), max(y,x) = (%d, %d)\n", g.MinY, g.MinX, g.MaxY, g.MaxX))
	b.WriteString("    ")
	for c := g.MinX; c <= g.MaxX; c++ {
		b.WriteString(fmt.Sprintf("%3d ", c))
	}
	b.WriteRune('\n')

	for r := g.MinY; r <= g.MaxY; r++ {
		b.WriteString(fmt.Sprintf("%3d ", r))
		for c := g.MinX; c <= g.MaxX; c++ {
			r, ok := g.Tiles[common.Point{X: c, Y: r}]
			if ok {
				b.WriteString(fmt.Sprintf("  %c ", r))
			} else {
				b.WriteString("  . ")
			}
		}
		b.WriteRune('\n')
	}

	return b.String()
}

func minmax(min, max, next int) (int, int) {
	if next < min {
		min = next
	}

	if next > max {
		max = next
	}

	return min, max
}

func (ds *Part1Solver) buildGrid(lines []string) *Grid {
	tiles := map[common.Point]rune{}
	minY, maxY := 0, 0
	minX, maxX := 0, 0

	var re = regexp.MustCompile(`x=(-?\d+), y=(-?\d+)`)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		matches := re.FindAllStringSubmatch(line, -1)
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

		distance := sensor.ManhattanDistance(beacon)

		//fmt.Println("sensor", sensor, "distance", distance)

		if sensor.Y <= ds.TargetY && sensor.Y+distance >= ds.TargetY {
			width := sensor.Y + distance - ds.TargetY

			for x := width * -1; x <= width; x++ {
				p := common.Point{X: sensor.X + x, Y: ds.TargetY}
				if _, ok := tiles[p]; !ok {
					tiles[p] = NoSensor
					minX, maxX = minmax(minX, maxX, p.X)
					minY, maxY = minmax(minY, maxY, p.Y)
				}
			}
		} else if sensor.Y >= ds.TargetY && sensor.Y-distance <= ds.TargetY {
			width := ds.TargetY - (sensor.Y - distance)
			//fmt.Println("sensor", sensor, "distance", distance, "width", width)

			for x := width * -1; x <= width; x++ {
				p := common.Point{X: sensor.X + x, Y: ds.TargetY}
				if _, ok := tiles[p]; !ok {
					tiles[p] = NoSensor
					minX, maxX = minmax(minX, maxX, p.X)
					minY, maxY = minmax(minY, maxY, p.Y)
				}
			}
		}
	}

	fmt.Printf("min(y,x) = (%d, %d), max(y,x) = (%d, %d)\n", minY, minX, maxY, maxX)

	return &Grid{Tiles: tiles, MinX: minX, MaxX: maxX, MinY: minY, MaxY: maxY}
}

type Part1Solver struct {
	TargetY int
}

func (ds *Part1Solver) Solve(scanner *bufio.Scanner) (common.Solution, error) {
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	grid := ds.buildGrid(lines)
	//fmt.Println(grid.String())

	var count int
	for x := grid.MinX; x <= grid.MaxX; x++ {
		if grid.Tiles[common.Point{X: x, Y: ds.TargetY}] == NoSensor {
			count++
		}
	}

	return common.Solution{IntValue: count}, nil
}

type Part2Solver struct{}

func (ds *Part2Solver) Solve(scanner *bufio.Scanner) (common.Solution, error) {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// TODO
	return common.Solution{}, nil
}
