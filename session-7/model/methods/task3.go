package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius

}
func main() {
	circle := Circle{Radius: 7}
	circle.Area()
	fmt.Printf("Circle radius : %f \nArea : %f \n", circle.Radius, circle.Area())
}
