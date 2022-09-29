package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxFloat(3, 7, 88.5, 5, 15))
}

func maxFloat(floats ...float64) float64 {
	max := math.Inf(-1)
	for _, v := range floats {
		if v > max {
			max = v
		}
	}
	return max
}
