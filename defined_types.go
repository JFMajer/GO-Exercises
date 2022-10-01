package main

import (
	"fmt"
)

type Liters float64
type Gallons float64

func main() {
	var carFuel Liters
	var busFuel Gallons
	carFuel = Liters(55)
	busFuel = Gallons(75)
	fmt.Println(carFuel, busFuel)
	fmt.Printf("75 Gallons equals to %.2f liters\n", carFuel.toGallons())
	fmt.Printf("55 Liters equals to %.2f gallons\n", busFuel.toLiters())
}

func (l Liters) toGallons() Gallons {
	return Gallons(l * 0.264)
}

func (g Gallons) toLiters() Liters {
	return Liters(g * 3.7854)
}
