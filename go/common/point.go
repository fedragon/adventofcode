package common

import (
	"math"
)

type Point struct {
	X, Y int
}

// ManhattanDistance returns the distance between two points, computed as |x1-x2| + |y1-y2|
func (p Point) ManhattanDistance(t Point) int {
	return int(math.Abs(float64(p.X-t.X))) + int(math.Abs(float64(p.Y-t.Y)))
}
