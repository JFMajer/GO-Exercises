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

	//var myArray [5]int
	//var mySlice []int
	//anotherSlice := make([]string, 9)
	//yetAnotherSlice := []int{9,8,7}
	yetAnotherSlice := []int{9, 8, 7}
	yetAnotherSlice = append(yetAnotherSlice, 5)
	fmt.Println(yetAnotherSlice)

}
