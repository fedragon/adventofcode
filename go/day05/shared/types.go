package shared

import (
	"fmt"
	"sort"
)

type Crane struct {
	Stacks map[int]*Stack
	Moves  []Move
}

func (c *Crane) AddCrate(index int, elem string) {
	stack, ok := c.Stacks[index]
	if !ok {
		stack = &Stack{}
		c.Stacks[index] = stack
	}

	stack.Push(elem)
}

func (c *Crane) AddMove(move Move) {
	c.Moves = append(c.Moves, move)
}

func (c *Crane) ExecutePlan(mover Mover) error {
	for _, m := range c.Moves {
		from, ok := c.Stacks[m.From]
		if !ok {
			return fmt.Errorf("source stack not found: %v", m.From)
		}
		to, ok := c.Stacks[m.To]
		if !ok {
			return fmt.Errorf("target stack not found: %v", m.To)
		}

		mover.Move(from, to, m.Count)
	}

	return nil
}

func (c *Crane) Finalize() {
	for _, stack := range c.Stacks {
		stack.Reverse()
	}
}

func (c *Crane) TopCrates() []string {
	keys := make([]int, 0)
	for key := range c.Stacks {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	crates := make([]string, 0)
	for _, key := range keys {
		crates = append(crates, c.Stacks[key].Peek())
	}

	return crates
}

func (c *Crane) PrintStacks() {
	for k, v := range c.Stacks {
		fmt.Println(k, ":", v)
	}
	fmt.Println("")
}

type Mover interface {
	Move(from *Stack, to *Stack, amount int)
}

type CrateMover9000 struct{}

// Move moves multiple elements from one stack to another, one by one: doing so, their order is inverted
func (c *CrateMover9000) Move(from *Stack, to *Stack, count int) {
	for i := 0; i < count; i++ {
		to.Push(from.Pop())
	}
}

type CrateMover9001 struct{}

// Move moves multiple elements from one stack to another, preserving their order
func (c *CrateMover9001) Move(from *Stack, to *Stack, count int) {
	to.PushN(from.PopN(count)...)
}

type Move struct {
	Count int
	From  int
	To    int
}
