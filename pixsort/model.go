package pixsort

import "math"

type Point struct {
	X, Y int
}

func (a Point) DistanceTo(b Point) float64 {
	dX := float64(a.X - b.X)
	dY := float64(a.Y - b.Y)

	return math.Pow(dX, 2) + math.Pow(dY, 2)
}
