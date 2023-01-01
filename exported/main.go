package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 100000, FullTime: true},
	{FirstName: "Lucy", LastName: "Jones", Salary: 50000, FullTime: false},
	{FirstName: "Alan", LastName: "Anderson", Salary: 25000, FullTime: false},
	{FirstName: "Margaret", LastName: "Carter", Salary: 120000, FullTime: true},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	log.Println(myStaff.All())

	log.Println(myStaff.Overpaid())
	log.Println(myStaff.Underpaid())

}
