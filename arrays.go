package main

import (
	"fmt"
)

func main() {
	ints := [3]int{1, 2}
	fmt.Println(ints)

	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	for i, v := range notes {
		fmt.Println(i, v)
	}
}
