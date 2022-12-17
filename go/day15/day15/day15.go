package day15

import (
	"bufio"
	"fmt"
	"github.com/fedragon/adventofcode/common"
	"regexp"
	"runtime"
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
	var areas []Area

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

		areas = append(areas, Area{Sensor: sensor, Distance: sensor.ManhattanDistance(&beacon)})
	}

	maxWorkers := runtime.GOMAXPROCS(-1)
	jobs := make(chan minMax, maxWorkers)
	result := make(chan common.Point, 1)

	for i := 0; i < maxWorkers; i++ {
		go findBeacon(i, areas, jobs, result)
	}

	const step = 2000
	i, j := 0, common.Min(ds.Max.X, step)
	var start, end common.Point
	for {
		start = common.Point{X: i, Y: i}
		end = common.Point{X: j, Y: j}

		if end.Compare(ds.Max) == 1 {
			break
		}

		jobs <- minMax{Min: start, Max: end}

		i += step
		j += step
	}
	close(jobs)

	beacon := <-result

	tuningFrequency := beacon.X*4000000 + beacon.Y
	return common.Solution{IntValue: tuningFrequency}, nil
}

type minMax struct {
	Min, Max common.Point
}

func findBeacon(id int, areas []Area, jobs <-chan minMax, result chan<- common.Point) {
	for job := range jobs {
		fmt.Printf("[%v] now checking %v -> %v\n", id, job.Min, job.Max)
		for x := job.Min.X; x <= job.Max.X; x++ {
			for y := job.Min.Y; y <= job.Max.Y; y++ {
				p := common.Point{X: x, Y: y}

				none := false
				for _, a := range areas {
					none = a.Contains(&p) || none
				}

				if !none {
					fmt.Printf("[%v] FOUND beacon at: %v\n", id, p)
					result <- p
					return
				}
			}
		}

		//fmt.Printf("[%v] found no beacon for: %v\n", id, job)
	}
}

type Area struct {
	Sensor   common.Point
	Distance int
}

func (s *Area) Contains(p *common.Point) bool {
	return s.Sensor.ManhattanDistance(p) <= s.Distance
}
