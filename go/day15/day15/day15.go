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

var re = regexp.MustCompile(`x=(-?\d+), y=(-?\d+)`)

type Grid struct {
	Min, Max *common.Point
	Tiles    map[common.Point]rune
}

func (g *Grid) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("min(x,y) = (%d, %d), max(x,y) = (%d, %d)\n", g.Min.X, g.Min.Y, g.Max.X, g.Max.Y))
	b.WriteString("    ")
	for c := g.Min.X; c <= g.Max.X; c++ {
		b.WriteString(fmt.Sprintf("%3d ", c))
	}
	b.WriteRune('\n')

	for r := g.Min.Y; r <= g.Max.Y; r++ {
		b.WriteString(fmt.Sprintf("%3d ", r))
		for c := g.Min.X; c <= g.Max.X; c++ {
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

		distance := sensor.ManhattanDistance(&beacon)

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

	fmt.Printf("min(x,y) = (%d, %d), max(x,y) = (%d, %d)\n", minX, minY, maxX, maxY)

	return &Grid{Tiles: tiles, Min: &common.Point{X: minX, Y: minY}, Max: &common.Point{X: maxX, Y: maxY}}
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

	var count int
	for x := grid.Min.X; x <= grid.Max.X; x++ {
		if grid.Tiles[common.Point{X: x, Y: ds.TargetY}] == NoSensor {
			count++
		}
	}

	return common.Solution{IntValue: count}, nil
}

type Part2Solver struct {
	Min, Max *common.Point
}

func (ds *Part2Solver) Solve(scanner *bufio.Scanner) (common.Solution, error) {
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// TODO
	return common.Solution{}, nil
}
