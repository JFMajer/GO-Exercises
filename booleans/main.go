package main

import "fmt"

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {
	apples := 18
	oranges := 9

	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)
	fmt.Println(apples < oranges)
	fmt.Println(apples > oranges)

	fmt.Printf("%d >= %d, %t\n", 3, 4, 3 > 4)

	jack := Employee{
		Name:     "Jack Smith",
		Age:      27,
		Salary:   40000,
		FullTime: false,
	}

	jill := Employee{
		Name:     "Jill Jones",
		Age:      33,
		Salary:   60000,
		FullTime: true,
	}

	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, x := range employees {
		if x.Age > 30 {
			fmt.Printf("%s is older than 30\n", x.Name)
		} else {
			fmt.Printf("%s is not older than 30\n", x.Name)
		}
	}
}
