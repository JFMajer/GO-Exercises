package main

import (
	"fmt"
	"sort"
)

func main() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "fish")

	fmt.Println(animals)

	for _, x := range animals {
		fmt.Println(x)
	}

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println(animals[0])

	fmt.Println(animals[:1])
	fmt.Println(animals[:2])
	fmt.Println(animals[:3])
	fmt.Println(animals[:4])

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))
	for i, x := range animals {
		fmt.Println(i, x)
	}

	arr := deleteFromSlice(animals, 2)
	for i, x := range arr {
		fmt.Println(i, x)
	}

}

func deleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]

	return a
}
