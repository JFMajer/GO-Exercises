package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Printf("Number of passengers: %d\n", v.NumberOfPassengers)
	fmt.Printf("Number of wheels: %d\n", v.NumberOfWheels)
}

func (c Car) show() {
	fmt.Printf("Make: %s\n", c.Make)
	fmt.Printf("Model: %s\n", c.Model)
	fmt.Printf("Year: %d\n", c.Year)
	fmt.Printf("Is electric: %t\n", c.IsElectric)
	fmt.Printf("Is hybrid: %t\n", c.IsHybrid)
	c.Vehicle.showDetails()
}

func main() {
	car := Car{
		Make:       "Tesla",
		Model:      "Model 3",
		Year:       2019,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle: Vehicle{
			NumberOfWheels:     4,
			NumberOfPassengers: 5,
		},
	}
	car.show()

	suv := Vehicle{
		NumberOfWheels:     4,
		NumberOfPassengers: 7,
	}
	suv.showDetails()
	volvoXC90 := Car{
		Make:       "Volvo",
		Model:      "XC90",
		Year:       2019,
		IsElectric: false,
		IsHybrid:   false,
		Vehicle:    suv,
	}
	volvoXC90.show()

	carWithNoWheels := Car{
		Make:       "Tesla",
		Model:      "Model 3",
		Year:       2019,
		IsElectric: true,
		IsHybrid:   false,
	}
	carWithNoWheels.show()
	volvoXC90.Vehicle.NumberOfWheels = 9
	volvoXC90.show()

}
