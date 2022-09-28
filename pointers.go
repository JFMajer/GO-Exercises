package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt)) //*int
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat)) //*float64

	var someInt int
	var someIntPointer *int
	someIntPointer = &someInt
	fmt.Println(someIntPointer)
	fmt.Println(*someIntPointer) //0

	var someFloat float64 = 2.4353
	someFloatPointer := &someFloat
	fmt.Println(someFloatPointer)
	fmt.Println(*someFloatPointer) //2.4353

	var yetAnotherInt int = 3
	fmt.Printf("Value of yetAnotherInt equals to %v\n", yetAnotherInt) //3
	yetAnotherIntPointer := &yetAnotherInt
	*yetAnotherIntPointer = 5
	fmt.Printf("Value of yetAnotherInt equals to %v now\n", yetAnotherInt) //5

	double(someFloatPointer)
	fmt.Println(*someFloatPointer)
}

func double(point *float64) {
	*point *= 2
}
