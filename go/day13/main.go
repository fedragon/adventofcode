package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fedragon/adventofcode/common"
	"os"
)

type Outcome int

const (
	RightOrder Outcome = iota - 1
	CantSay
	WrongOrder
)

type Value interface {
	Compare(Value) Outcome
}

type Int int

func (i Int) Compare(v Value) Outcome {
	switch x := v.(type) {
	case Int:
		a := int(i)
		b := int(x)

		switch {
		case a < b:
			return RightOrder
		case a > b:
			return WrongOrder
		}
		return CantSay

	case List:
		return List{i}.Compare(v)
	}

	panic("unexpected type")
}

type List []Value

func (l List) Compare(v Value) Outcome {
	var other List

	switch x := v.(type) {
	case Int:
		other = List{x}
	case List:
		other = x
	}

	if len(l) == 0 && len(other) > 0 {
		return RightOrder
	}

	if len(l) > 0 && len(other) == 0 {
		return WrongOrder
	}

	for idx, value := range l {
		if idx == len(other) {
			break
		}

		if outcome := value.Compare(other[idx]); outcome != 0 {
			return outcome
		}
	}

	// still can't say anything: decide by length
	return Int(len(l)).Compare(Int(len(other)))
}

type Part1Solver struct{}

func (ds *Part1Solver) Solve(scanner *bufio.Scanner) (int, error) {
	var left string
	var sum int

	pairIndex := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		if len(line) == 0 {
			continue
		}

		if len(left) == 0 {
			left = line
			continue
		}

		outcome := parseJSON(left).Compare(parseJSON(line))
		fmt.Println("index", pairIndex, "outcome", outcome)
		if outcome < WrongOrder {
			sum += pairIndex
		}
		pairIndex++
		left = ""
	}

	return sum, nil
}

func main() {
	f := common.Must(os.Open("../data/day13"))
	defer f.Close()

	part1 := Part1Solver{}
	solution, err := part1.Solve(bufio.NewScanner(f))
	if err != nil {
		panic(err)
	}

	fmt.Println("solution for part 1", solution)

	//common.Must(f.Seek(0, 0))
	//
	//part2 := PartXSolver{}
	//solution, err = part2.Solve(bufio.NewScanner(f))
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("solution for part 2", solution)
}

func parseJSON(line string) List {
	var parsed []any

	if err := json.Unmarshal([]byte(line), &parsed); err != nil {
		panic(err)
	}

	return traverse(parsed, nil)
}

func traverse(xs []any, acc List) List {
	if len(xs) == 0 {
		return acc
	}

	for _, x := range xs {
		switch e := x.(type) {
		case float64:
			acc = append(acc, Int(e))
		case []any:
			acc = append(acc, traverse(e, nil))
		default:
			panic("unexpected type")
		}
	}

	return acc
}
