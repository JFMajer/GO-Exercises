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
	fmt.Printf("Random number is %v\n", ran)

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
	fmt.Println(guess)
}
