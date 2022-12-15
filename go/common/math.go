package common

import (
	"fmt"
	"math"
)

func Abs(i int) int {
	return int(math.Abs(float64(i)))
}

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

type Range struct {
	Start, End int
}

func (r *Range) String() string {
	return fmt.Sprintf("%v ... %v", r.Start, r.End)
}

func (r *Range) Overlaps(a *Range) bool {
	if r.End < a.Start {
		return false
	} else if a.End < r.Start {
		return false
	}

	return true
}

func (r *Range) FullyContains(a *Range) bool {
	return r.Start <= a.Start && r.End >= a.End
}

func (r *Range) Merge(a *Range) *Range {
	if !r.Overlaps(a) {
		return nil
	}

	res := &Range{}
	if r.Start < a.Start {
		res.Start = r.Start
	} else {
		res.Start = a.Start
	}

	if r.End > a.End {
		res.End = r.End
	} else {
		res.End = a.End
	}

	return res
}

type Ranges []Range

func (r Ranges) Len() int {
	return len(r)
}

func (r Ranges) Less(i, j int) bool {
	return r[i].Start < r[j].Start
}

func (r Ranges) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
