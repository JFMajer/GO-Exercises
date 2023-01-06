package main

import (
	"fmt"
	"strings"
)

func main() {
	courseName := "Learn Go for Beginners Crash Course"

	fmt.Println(courseName[0])
	fmt.Println(string(courseName[0]))

	for i := 0; i < len(courseName); i++ {
		fmt.Print(string(courseName[i]))
	}

	fmt.Println()
	fmt.Println("=======================================")

	var mySlice []string
	mySlice = append(mySlice, "one")
	mySlice = append(mySlice, "two")
	mySlice = append(mySlice, "three")
	fmt.Printf("My slice has length of %d elements\n", len(mySlice))
	fmt.Println(mySlice)
	for i, e := range mySlice {
		fmt.Printf("Index: %d, Element: %s\n", i, e)
	}

	fmt.Println("=======================================")
	courses := []string{
		"Learn Go for Beginners Crash Course",
		"Learn Java for Beginners Crash Course",
		"Learn C for Beginners Crash Course",
	}

	for _, ele := range courses {
		if strings.Contains(ele, "Go") {
			fmt.Println("This one contains Go and index is: ", strings.Index(ele, "Go"))
			fmt.Println(ele)
		} else {
			fmt.Println("This one doesn't")
			fmt.Println(ele)
		}
	}

	newString := "Go is a great programming language. Go for it! But Python is also not bad."
	fmt.Println(strings.HasPrefix(newString, "Go"))
	fmt.Println(strings.HasPrefix(newString, "Python"))
	fmt.Println(strings.HasPrefix(newString, "go"))
	fmt.Println(strings.HasSuffix(newString, "go"))
	fmt.Println(strings.HasSuffix(newString, "bad."))
	fmt.Println(strings.Count(newString, "Go"))
	fmt.Println(strings.Count(newString, "go"))
	fmt.Println(strings.Count(strings.ToLower(newString), "go"))

	badEmail := " bad@email.com "
	goodEmail := strings.TrimSpace(badEmail)
	fmt.Println(badEmail)
	fmt.Println(goodEmail)

}
