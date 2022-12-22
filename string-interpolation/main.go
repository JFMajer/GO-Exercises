package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func main() {
	name1 := getString("What is your name?")
	fmt.Printf("Your name is %s\n", name1)
	age1 := getNumber("How old are you?")
	fmt.Printf("You are %d years old\n", age1)

}

// function that grabs string from the user
func getString(q string) string {
	fmt.Println(q)
	fmt.Print("-> ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)
	return userInput
}

func getNumber(q string) int {
	fmt.Println(q)
	fmt.Print("-> ")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)
	return convertStrToInt(userInput)
}

func convertStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error during conversion")
		log.Fatal(err)
	}
	return num
}
