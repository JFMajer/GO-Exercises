package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	OwnsACat        bool
}

func main() {
	var user User

	user.UserName = getString("What is your name?")
	user.Age = getNumber("How old are you?")
	user.FavouriteNumber = getFloat("What is your favourite number?")

	fmt.Printf("Your name is %s and you are %d years old\n", user.UserName, user.Age)
	fmt.Printf("Your favourite number is %.2f.\n", user.FavouriteNumber)

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

// function that grabs number from user
func getNumber(q string) int {
	var valueToReturn int
	for {
		fmt.Println(q)
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		valueToReturn = convertStrToInt(userInput)
		// run loop again if conversion failed
		if valueToReturn != 0 {
			break
		}
	}
	return valueToReturn
}

// function that grabs float from user
func getFloat(q string) float64 {
	var valueToReturn float64
	for {
		fmt.Println(q)
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		valueToReturn = convertStrToFloat(userInput)
		// run loop again if conversion failed
		if valueToReturn != 0 {
			break
		}
	}
	return valueToReturn
}

// function that converts string to int
func convertStrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error during conversion, please provide valid positive number")
		//log.Fatal(err)
	}
	return num
}

// function that converts string to float64
func convertStrToFloat(s string) float64 {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Error during conversion, please provide valid number")
		//log.Fatal(err)
	}
	return num
}
