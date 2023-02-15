package helper

import "math"

func Distance(x1, y1, x2, y2 float64) float64 {
	x, y := x2-x1, y2-y1
	return math.Sqrt(x*x + y*y)
}

func Length(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}
