//guessing game written in go

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 100
	numToGuess := rand.Intn(max-min) + min

	fmt.Println("Guess a number between 0 and 100, type q to quit")

	for {
		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "q" {
			fmt.Printf("Game Over, the number was %d\n", numToGuess)
			break
		}

		guess, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Error during conversion")
		}

		if guess < 0 || guess > 100 {
			fmt.Println("Please provide a number between 0 and 100")
		} else if guess < numToGuess {
			fmt.Println("Number to guess is higher")
		} else if guess > numToGuess {
			fmt.Println("Number to guess is lower")
		} else {
			fmt.Println("Correct!")
			break
		}

	}

}
