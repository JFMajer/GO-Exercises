// average2 calculates the average of several numbers provided as command line arguments
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var sum float64
	for _, v := range arguments {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += f
	}
	fmt.Printf("Average equals to %.2f\n", sum/float64(len(arguments)))
}
