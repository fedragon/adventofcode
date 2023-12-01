package day15

import (
	"bufio"
	"context"
	"fmt"
	"github.com/fedragon/adventofcode/common"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

const (
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

type sensorArea struct {
	Sensor   common.Point
	Distance int
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
		tiles[sensor] = 'S'

		minX, maxX = minmax(minX, maxX, sensor.X)
		minY, maxY = minmax(minY, maxY, sensor.Y)

		beacon := common.Point{
			X: common.Must(strconv.Atoi(matches[1][1])),
			Y: common.Must(strconv.Atoi(matches[1][2])),
		}
		tiles[beacon] = 'B'

		minX, maxX = minmax(minX, maxX, beacon.X)
		minY, maxY = minmax(minY, maxY, beacon.Y)

		distance := sensor.ManhattanDistance(&beacon)

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
	var areas []sensorArea

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		matches := re.FindAllStringSubmatch(line, -1)
		sensor := common.Point{
			X: common.Must(strconv.Atoi(matches[0][1])),
			Y: common.Must(strconv.Atoi(matches[0][2])),
		}

		beacon := common.Point{
			X: common.Must(strconv.Atoi(matches[1][1])),
			Y: common.Must(strconv.Atoi(matches[1][2])),
		}

		areas = append(areas, sensorArea{Sensor: sensor, Distance: sensor.ManhattanDistance(&beacon)})
	}

	ctx, cancel := context.WithCancel(context.Background())
	maxWorkers := runtime.GOMAXPROCS(-1)
	rows := make(chan int, maxWorkers)
	result := make(chan common.Point, 1)

	for i := 0; i < maxWorkers; i++ {
		go findBeacon(ctx, i, areas, rows, result)
	}

	go func(jobs chan<- int) {
		for y := ds.Min.Y; y <= ds.Max.Y; y++ {
			jobs <- y
		}
		close(jobs)
	}(rows)

	beacon := <-result
	cancel()

	tuningFrequency := beacon.X*4000000 + beacon.Y
	return common.Solution{IntValue: tuningFrequency}, nil
}

func findBeacon(ctx context.Context, id int, areas []sensorArea, rows <-chan int, result chan<- common.Point) {
	for {
		select {
		case <-ctx.Done():
			return
		case row := <-rows:
			fmt.Printf("[%v] now checking row %v\n", id, row)
			var ranges common.Ranges
			for _, a := range areas {
				width := a.Distance - common.Abs(a.Sensor.Y-row)
				if width < 0 {
					continue
				}

				ranges = append(ranges, common.Range{Start: a.Sensor.X - width, End: a.Sensor.X + width})
			}

			var start common.Range
			var found bool
			if len(ranges) > 0 {
				sort.Sort(ranges)
				start = ranges[0]

				for i := 1; i < len(ranges); i++ {
					current := ranges[i]
					res := start.Merge(&current)

					if res == nil {
						found = true
						break
					}

					start = *res
				}
			}

			if found {
				result <- common.Point{X: start.End + 1, Y: row}
				return
			}
		}
	}
}
