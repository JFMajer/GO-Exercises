// go mod init scope
package main

import (
	"fmt"
	"scope/package1"
)

var one string = "three"
var two string = "i'm global variable"

func main() {
	var one string = "One"

	fmt.Println(one)
	myFunc()

	// variable shadowing
	two := "i'm block level variable"
	fmt.Println(two) // will print block level variable

	newString := package1.PublicVar
	fmt.Printf("From package1: %s\n", newString)
}

func myFunc() {
	// var one string = "Two"
	fmt.Println(one)
}
