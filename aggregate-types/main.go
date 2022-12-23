package main

import "fmt"

type Car struct {
	Luxury      bool
	HorsePower  int
	BucketSeats bool
	Make        string
	Model       string
	Year        int
}

func main() {
	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "chicken"

	fmt.Println(myStrings)

	for i, s := range myStrings {
		fmt.Println(i, s)
	}

	var myCar Car
	myCar.BucketSeats = true
	myCar.HorsePower = 320
	myCar.Make = "Audi"

	fmt.Println(myCar)

	mySecondCar := Car{
		Luxury:      true,
		HorsePower:  150,
		BucketSeats: false,
		Make:        "Mercedes",
		Model:       "A150",
		Year:        2014,
	}

	fmt.Println(mySecondCar)

}
