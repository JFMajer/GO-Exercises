package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	rand.Seed(time.Now().UnixNano())

	x := 1000
	for x >= 100 {
		//get a random number between 1 and 1001 and assign to variable x
		x = rand.Intn(1000) + 1
		//print the value of x
		if x < 100 {
			fmt.Printf("The value of x is %d and it is less than 100\n", x)
		} else {
			fmt.Printf("The value of x is %d and it is greater than 100\n", x)
		}
	}

}
