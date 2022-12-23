package main

import (
	"fmt"
	"log"
)

var myInt int

// var myInt16 int16
// var myInt32 int32
// var myInt64 int64
var myUint uint // usnigned integer, only positive values
var myFloat float32
var myFloat64 float64

func main() {
	myInt = -10
	myUint = 1
	myFloat = 10.1
	myFloat64 = 34784792.84728472

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := ""
	log.Println(myString)

	var myBool = true
	log.Println(myBool)

	if myBool {
		fmt.Println("true")
	}

	if myBool {
		fmt.Println(!myBool)
	}
}
