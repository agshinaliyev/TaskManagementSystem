package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {

	return r.Width * r.Height

}
func (r *Rectangle) Perimeter() float64 {

	return (r.Width + r.Height) * 2
}
func main() {

	r := Rectangle{
		Width:  10.5,
		Height: 5.0,
	}

	fmt.Printf("Width : %f , Height : %f \nArea: %f\nPerimeter : %f \n", r.Width, r.Height, r.Area(), r.Perimeter())
}
