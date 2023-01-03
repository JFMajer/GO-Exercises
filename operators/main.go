package main

import (
	"fmt"
	"math"
)

const (
	PI = 3.14
)

func main() {
	answer := 7 + 3*4
	fmt.Printf("answer is %d\n", answer)

	answer = (7 + 3) * 4
	fmt.Printf("answer is %d\n", answer)
	var radius float64
	radius = 12.0
	area := PI * radius * radius
	fmt.Printf("area is %f\n", area)
	area = math.Pi * radius * radius
	fmt.Printf("area is %f\n", area)

	half := 3 / 2.0
	fmt.Println(half)
}
