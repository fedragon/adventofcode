package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`^([[:alpha:]]) (\d+)$`)

var trailFile *os.File

func init() {
	var err error
	trailFile, err = os.Create("trail.txt")
	if err != nil {
		panic(err)
	}
}

type History map[int]map[int]rune

type Move struct {
	Direction string
	Steps     int
}

type Knot struct {
	Marker   rune
	Row, Col int
	Visited  map[string]struct{}
}

func (k *Knot) Follow(h *Knot) {
	if k.IsAdjacentTo(*h) {
		return
	}

	if h.Row < k.Row {
		k.Row--
	} else if h.Row > k.Row {
		k.Row++
	}

	if h.Col < k.Col {
		k.Col--
	} else if h.Col > k.Col {
		k.Col++
	}

	k.Visited[fmt.Sprintf("(%d,%d)", k.Row, k.Col)] = struct{}{}
}

func (k *Knot) IsAdjacentTo(o Knot) bool {
	return math.Abs(float64(k.Row-o.Row)) <= 1 && math.Abs(float64(k.Col-o.Col)) <= 1
}

func main() {
	f, err := os.Open("../data/day09")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var moves []Move
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			panic(fmt.Errorf("expected two matches on line '%v', but got %v", line, matches))
		}

		steps, _ := strconv.Atoi(matches[2])
		moves = append(moves, Move{Direction: matches[1], Steps: steps})

	}

	fmt.Println("visited (2 knots)", Count(Visit(moves)))

	knots := make([]*Knot, 10)
	for i := 0; i < 10; i++ {
		knots[i] = &Knot{rune('0' + i), 0, 0, map[string]struct{}{"(0,0)": {}}}
	}

	VisitN(knots, moves)

	fmt.Println("visited (10 knots)", len(knots[9].Visited))
}

func Visit(moves []Move) History {
	head := Knot{'H', 0, 0, nil}
	tail := Knot{'T', 0, 0, nil}
	history := History{0: {0: 'T'}}

	follow := func(tail, prev, next Knot) Knot {
		if !tail.IsAdjacentTo(next) {
			row, col := prev.Row, prev.Col
			if history[row] == nil {
				history[row] = make(map[int]rune)
			}
			history[row][col] = 'T'
			tail = prev
		}

		LogTrail(history, head)

		return tail
	}

	for _, m := range moves {
		switch m.Direction {
		case "U":
			for rem := 0; rem < m.Steps; rem++ {
				prev := head
				head.Row--
				tail = follow(tail, prev, head)

				LogTrail(history, head)
			}
		case "D":
			for rem := 0; rem < m.Steps; rem++ {
				prev := head
				head.Row++
				tail = follow(tail, prev, head)

				LogTrail(history, head)
			}
		case "L":
			for rem := 0; rem < m.Steps; rem++ {
				prev := head
				head.Col--
				tail = follow(tail, prev, head)

				LogTrail(history, head)
			}
		case "R":
			for rem := 0; rem < m.Steps; rem++ {
				prev := head
				head.Col++
				tail = follow(tail, prev, head)

				LogTrail(history, head)
			}
		default:
			panic(fmt.Errorf("unexpected direction: %v", m.Direction))
		}
	}

	return history
}

func VisitN(knots []*Knot, moves []Move) {
	for _, m := range moves {
		switch m.Direction {
		case "U":
			for rem := 0; rem < m.Steps; rem++ {
				knots[0].Row--

				for i := 1; i < len(knots); i++ {
					knots[i].Follow(knots[i-1])
				}
			}
		case "D":
			for rem := 0; rem < m.Steps; rem++ {
				knots[0].Row++

				for i := 1; i < len(knots); i++ {
					knots[i].Follow(knots[i-1])
				}
			}
		case "L":
			for rem := 0; rem < m.Steps; rem++ {
				knots[0].Col--

				for i := 1; i < len(knots); i++ {
					knots[i].Follow(knots[i-1])
				}
			}
		case "R":
			for rem := 0; rem < m.Steps; rem++ {
				knots[0].Col++

				for i := 1; i < len(knots); i++ {
					knots[i].Follow(knots[i-1])
				}
			}
		default:
			panic(fmt.Errorf("unexpected direction: %v", m.Direction))
		}
	}
}

func Count(history History) int {
	total := 0
	for _, row := range history {
		for _, value := range row {
			if value == 'T' {
				total++
			}
		}
	}
	return total
}

func LogTrail(history History, head Knot) {
	if _, ok := os.LookupEnv("LOG_TRAIL"); !ok {
		return
	}

	m := make(History)

	for r, row := range history {
		m[r] = make(map[int]rune)

		for c, v := range row {
			m[r][c] = v
		}
	}

	if m[head.Row] == nil {
		m[head.Row] = make(map[int]rune)
	}
	m[head.Row][head.Col] = 'H'

	var sortedKeys []int
	for k := range m {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)

	all := func(m map[int]rune) (int, int) {
		var keys []int
		for k := range m {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		min := keys[0]
		max := keys[len(keys)-1]

		if max < 0 {
			max = 0
		}

		if min > 0 {
			min = 0
		}

		return min, max
	}

	var min, max int
	for _, r := range sortedKeys {
		lmin, lmax := all(m[r])
		if lmin < min {
			min = lmin
		}

		if lmax > max {
			max = lmax
		}
	}

	var b strings.Builder
	var row []string
	for _, r := range sortedKeys {
		for c := min; c <= max; c++ {
			v, ok := m[r][c]
			if ok {
				row = append(row, string(v))
			} else {
				row = append(row, ".")
			}
		}

		b.WriteString(fmt.Sprintf("%2d: ", r))
		b.WriteString(strings.Join(row, " "))
		b.WriteRune('\n')
		row = nil
	}

	trailFile.WriteString(b.String())
	trailFile.WriteString("\n")
}
