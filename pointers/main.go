package main

import "fmt"

func main() {
	x := 10
	myFirstPointer := &x

	fmt.Printf("x is %d\n", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15
	fmt.Println("x is now", x)

	changeValueOfPointer(&x)
	fmt.Println("x is now", x)
}

func changeValueOfPointer(num *int) {
	*num = 25
}
