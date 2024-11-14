package main

import "fmt"

type Vehicle struct {
	Make  string
	Model string
	Year  int
}
type Car struct {
	Vehicle
	FuelType string
}

func main() {

	car := Car{
		Vehicle: Vehicle{Make: "Kia",
			Model: "K5",
			Year:  2022,
		},
		FuelType: "Gasoline",
	}

	fmt.Printf("Make : %s \nModel : %s \nYear : %d \nFuelType : %s",
		car.Make, car.Model, car.Year, car.FuelType)
}
