package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/JFMajer/keyboard"
)

func main() {
	//grade, err := getFloat()
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(grade)
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Printf("status is %s\n", status)
}

func getFloat() (float64, error) {
	fmt.Print("Please input your grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		log.Fatal(err)
	}
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return grade, nil
}
