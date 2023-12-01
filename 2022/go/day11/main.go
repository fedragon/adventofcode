package main

import (
	"bufio"
	"fmt"
	stdlog "log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	logsEnabled bool
)

func init() {
	_, logsEnabled = os.LookupEnv("LOGS_ENABLED")
}

type WorryLevel int

type Part struct {
	Rounds  int
	Monkeys []*Monkey
	Relieve func(WorryLevel) WorryLevel
}

type SortedMonkeys []*Monkey

func (m SortedMonkeys) Len() int {
	return len(m)
}

func (m SortedMonkeys) Less(i, j int) bool {
	return m[i].Inspections < m[j].Inspections
}

func (m SortedMonkeys) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type Monkey struct {
	ID          string
	Inspections int

	Items         []WorryLevel
	IncreaseWorry func(WorryLevel) WorryLevel
	TestOperand   int
	WhenTrue      func([]*Monkey, WorryLevel)
	WhenFalse     func([]*Monkey, WorryLevel)
}

func (m *Monkey) Catches(item WorryLevel) {
	m.Items = append(m.Items, item)
}

func (m *Monkey) Acts(monkeys []*Monkey, relieve func(WorryLevel) WorryLevel) {
	inspect := func(item WorryLevel) WorryLevel {
		logf("[Monkey %s] inspects an item with a worry level of %d\n", m.ID, item)
		level := m.IncreaseWorry(item)
		logf("[Monkey %s] Worry level is now %d\n", m.ID, level)

		level = relieve(level)
		logf("[Monkey %s] Gets bored with item. Worry level is now %d\n", m.ID, level)

		return level
	}

	for _, item := range m.Items {
		level := inspect(item)
		m.Inspections++

		if level%WorryLevel(m.TestOperand) == 0 {
			logf("[Monkey %s] Worry level %d is divisible by %d\n", m.ID, level, m.TestOperand)
			m.WhenTrue(monkeys, level)
		} else {
			logf("[Monkey %s] Worry level %d is not divisible by %d\n", m.ID, level, m.TestOperand)
			m.WhenFalse(monkeys, level)
		}
	}
	m.Items = nil
}

var rexs = []*regexp.Regexp{
	regexp.MustCompile(`^Monkey (\d+):$`),
	regexp.MustCompile(`^\s+Starting items: ([0-9, ]+)$`),
	regexp.MustCompile(`^\s+Operation: new = old ([+*]) (\w+)$`),
	regexp.MustCompile(`^\s+Test: divisible by (\d+)$`),
	regexp.MustCompile(`^\s+If true: throw to monkey (\d+)$`),
	regexp.MustCompile(`^\s+If false: throw to monkey (\d+)$`),
}

func main() {
	f, err := os.Open("../data/day11")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	startingMonkeys := make([]*Monkey, 0)

	scanner := bufio.NewScanner(f)
	var batch []string
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			startingMonkeys = parse(startingMonkeys, batch)
			batch = nil
		} else {
			batch = append(batch, line)
		}
	}
	startingMonkeys = parse(startingMonkeys, batch)

	otherMonkeys := make([]*Monkey, 0, len(startingMonkeys))
	for _, m := range startingMonkeys {
		m := *m
		otherMonkeys = append(otherMonkeys, &m)
	}

	relief := 1
	for _, m := range startingMonkeys {
		relief *= m.TestOperand
	}

	parts := []Part{
		{
			Rounds:  20,
			Monkeys: startingMonkeys,
			Relieve: func(level WorryLevel) WorryLevel {
				return WorryLevel(int(math.Floor(float64(level) / 3)))
			},
		},
		{
			Rounds:  10000,
			Monkeys: otherMonkeys,
			Relieve: func(level WorryLevel) WorryLevel {
				// TIL: using the product of all TestArgument as modulo operand doesn't change the properties of its result!
				return level % WorryLevel(relief)
			},
		},
	}

	for x, part := range parts {
		fmt.Println("Part", x+1)
		monkeys := part.Monkeys

		for r := 0; r < part.Rounds; r++ {
			for _, m := range monkeys {
				m.Acts(monkeys, part.Relieve)
			}

			if x == 0 || r%1000 == 0 {
				fmt.Printf("Round %d\n", r)
				for _, m := range monkeys {
					fmt.Printf("Monkey %s: %v\n", m.ID, m.Items)
				}
				fmt.Println("---")
			}
		}

		sort.Sort(sort.Reverse(SortedMonkeys(monkeys)))

		for _, m := range monkeys {
			fmt.Printf("Monkey %s inspected items %d times.\n", m.ID, m.Inspections)
		}

		fmt.Println("monkey business level", monkeys[0].Inspections*monkeys[1].Inspections)
	}
}

func parse(monkeys []*Monkey, batch []string) []*Monkey {
	var m = Monkey{}

	for i, l := range batch {
		matches := rexs[i].FindStringSubmatch(l)
		switch i {
		case 0:
			m.ID = matches[1]
		case 1:
			items := strings.Split(matches[1], ", ")
			for _, it := range items {
				x, _ := strconv.Atoi(it)
				m.Catches(WorryLevel(x))
			}
		case 2:
			operation := matches[1]
			operand := matches[2]

			if operation == "*" {
				if operand == "old" {
					m.IncreaseWorry = func(level WorryLevel) WorryLevel {
						return level * level
					}
				} else {
					value, _ := strconv.Atoi(operand)
					m.IncreaseWorry = func(level WorryLevel) WorryLevel {
						return level * WorryLevel(value)
					}
				}
			} else if operation == "+" {
				if operand == "old" {
					m.IncreaseWorry = func(level WorryLevel) WorryLevel {
						return level + level
					}
				} else {
					value, _ := strconv.Atoi(operand)
					m.IncreaseWorry = func(level WorryLevel) WorryLevel {
						return level + WorryLevel(value)
					}
				}
			} else {
				panic(fmt.Sprintf("unknown operation: %s", operation))
			}
		case 3:
			m.TestOperand, _ = strconv.Atoi(matches[1])
		case 4:
			target, _ := strconv.Atoi(matches[1])
			m.WhenTrue = func(monkeys []*Monkey, item WorryLevel) {
				monkeys[target].Catches(item)
			}
		case 5:
			target, _ := strconv.Atoi(matches[1])
			m.WhenFalse = func(monkeys []*Monkey, item WorryLevel) {
				monkeys[target].Catches(item)
			}
		}
	}

	return append(monkeys, &m)
}

func logf(s string, args ...any) {
	if logsEnabled {
		stdlog.Printf(s, args...)
	}
}
