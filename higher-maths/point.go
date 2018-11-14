package maths

import (
	"fmt"
	"math"
)

// Point eliptic curve point
type Point struct {
	x int32
	y int32
	a int32
	b int32
}

// NewPoint create a new Point
func NewPoint(x int32, y int32, a int32, b int32) (Point, error) {
	want := int32(math.Pow(float64(y), 2))
	got := int32(math.Pow(float64(x), 3)) + (a * x) + b
	if want != got {
		return Point{}, fmt.Errorf("Point (%d, %d) is not on the curve", x, y)
	}
	return Point{
		x: x,
		y: y,
		a: a,
		b: b,
	}, nil
}
