package main

import "fmt"

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal) Says() {
	fmt.Printf("A %s says %s\n", a.Name, a.Sound)
}

func (a *Animal) NumOfLegs() {
	fmt.Printf("Animal %s has %d legs\n", a.Name, a.NumberOfLegs)
}

func main() {
	fmt.Println(sumMany(1, 2, 3, 10))
	fmt.Println(sumMany(1, 5, 150, 67, 45, 33, 2, 1))

	var dog Animal
	dog.Name = "Reksio"
	dog.Sound = "Woof Woof"
	dog.NumberOfLegs = 4

	dog.Says()

	cat := Animal{
		Name:         "Kropek",
		Sound:        "Miau",
		NumberOfLegs: 4,
	}

	cat.Says()
	dog.NumOfLegs()
	cat.NumOfLegs()

}

func sumMany(nums ...int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}
