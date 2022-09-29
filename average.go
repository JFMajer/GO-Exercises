// average calculates average of several numbers
package main

import (
	"fmt"
	"log"

	"github.com/JFMajer/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Printf("Average is %.2f\n", sum/float64(len(numbers)))
}
