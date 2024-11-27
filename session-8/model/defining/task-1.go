package defining

import "math"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}
type Rectangle struct {
	Length, Width float64
}

func (a Circle) Area() float64 {

	return math.Pi * a.Radius * a.Radius
}

// why cannot be pointer mentioned as a receiver?

func (r Rectangle) Area() float64 {

	return r.Length * r.Width
}
