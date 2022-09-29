package main

import (
	"fmt"
	"log"

	"github.com/JFMajer/datafile"
)

func main() {
	votes := make(map[string]int)
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range lines {
		if _, ok := votes[v]; ok {
			votes[v] += 1
		} else {
			votes[v] = 1
		}
		fmt.Println(v)
	}
	fmt.Println(votes)
}
