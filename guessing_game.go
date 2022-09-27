package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand" //import path is math/rand, package name is rand
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	ran := rand.Intn(100) + 1
	//fmt.Printf("Random number is %v\n", ran)

	for i := 9; i >= 0; i-- {
		fmt.Print("Please input your guess: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			log.Fatal(err)
		}
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < ran {
			fmt.Println("Your guess was too low")
			fmt.Printf("You have %v guesses left\n", i)
		} else if guess > ran {
			fmt.Println("Your guess was too high")
			fmt.Printf("You have %v guesses left\n", i)
		} else {
			fmt.Println("You guessed correctly!")
			fmt.Printf("It took you %v guesses\n", 10-i)
			break
		}
	}

}
