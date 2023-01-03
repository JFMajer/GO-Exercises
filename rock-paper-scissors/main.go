package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(3)

	reader := bufio.NewReader(os.Stdin)

	clearScreen()
	fmt.Printf("Computer picked %d\n", computerValue)
	fmt.Print("Please enter rock, paper or scissors -> ")
	playerChoice, _ = reader.ReadString('\n')
	// to lower case
	playerChoice = strings.ToLower(strings.TrimSpace(playerChoice))

	switch playerChoice {
	case "rock":
		playerValue = ROCK
	case "paper":
		playerValue = PAPER
	case "scissors":
		playerValue = SCISSORS
	default:
		fmt.Println("Invalid choice")
	}

	computerChoice := ""
	switch computerValue {
	case 1:
		computerChoice = "paper"
	case 2:
		computerChoice = "scissors"
	case 0:
		computerChoice = "rock"
	}

	fmt.Println()
	fmt.Printf("Player picked %d which is %s, computer picked %d which means %s\n", playerValue, playerChoice, computerValue, computerChoice)

}

func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
