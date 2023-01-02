package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i is %d\n", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("   j is %d\n", j)
		}
	}
}
