package main

import (
	"fmt"

	"github.com/JFMajer/magazine"
)

func main() {
	sub1 := magazine.CreateSub("Jakub")
	fmt.Println(sub1)
	magazine.ShowSub(sub1)

	var empl1 magazine.Employee
	empl1.Name = "Bob"
	empl1.Salary = 65000
	fmt.Println(empl1)
	empl1.City = "Miami"
	fmt.Println(empl1)
	empl1.Street = "Sunshine Blvd."
	empl1.Country = "USA"
	empl1.PostalCode = "12345"
	fmt.Println(empl1)

	var empl2 magazine.Employee
	fmt.Println(empl2)
	empl2.Name = "Patrick"
	fmt.Println(empl2)
	empl2.City = "New York"
	fmt.Println(empl2)
}
