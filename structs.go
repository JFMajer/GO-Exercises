package main

import (
	"fmt"
)

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func showSub(sub subscriber) {
	if sub.active {
		fmt.Printf("%s has status active with rate %.2f\n", sub.name, sub.rate)
	} else {
		fmt.Printf("%s has status inactive with rate %.2f\n", sub.name, sub.rate)
	}
}

func createSub(name string) subscriber {
	var newSub subscriber
	newSub.active = true
	newSub.rate = 4.99
	newSub.name = name

	return newSub
}

func main() {
	var myStruct struct {
		number float64
		word   string
		toggle bool
	}

	fmt.Printf("%#v\n", myStruct)

	myStruct.number = 74.8
	fmt.Printf("%.2f\n", myStruct.number)

	// var subscriber struct {
	// 	name   string
	// 	rate   float64
	// 	active bool
	// }

	// subscriber.name = "Bob Bobbington"
	// subscriber.rate = 4.99
	// subscriber.active = true

	// fmt.Println(subscriber)

	var porsche car
	porsche.name = "911"
	porsche.topSpeed = 250
	fmt.Println(porsche)

	var subscriber1 subscriber
	subscriber1.name = "Bob"
	subscriber1.active = true
	subscriber1.rate = 4.99

	var subscriber2 subscriber
	subscriber2.name = "Angie"
	subscriber2.active = false
	subscriber2.rate = 5.99

	fmt.Println(subscriber1)
	showSub(subscriber1)
	showSub(subscriber2)

	subscriber3 := createSub("Teddy")
	showSub(subscriber3)

}
