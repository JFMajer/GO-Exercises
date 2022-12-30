package main

import "fmt"

func main() {
	age := 10
	name := "Jack"
	rightHanded := true

	fmt.Printf("%s is %d years old. Right handed: %t", name, age, rightHanded)
	fmt.Println()

	ageInTenYears := age + 10
	fmt.Printf("In ten years, %s will be %d\n", name, ageInTenYears)
	isATeenager := age >= 13
	fmt.Printf("%s is a teenager? %t\n", name, isATeenager)
}
