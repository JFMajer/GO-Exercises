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

	playerChoice := ""
	playerValue := -1

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	playerScore := 0
	computerScore := 0

	for {
		rand.Seed(time.Now().UnixNano())
		computerValue := rand.Intn(3)
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

		if playerValue == computerValue {
			fmt.Println("It's a draw!!")
		} else if playerValue == ROCK {
			switch computerValue {
			case 1:
				computerScore += 1
			case 2:
				playerScore += 1
			}
		} else if playerValue == PAPER {
			switch computerValue {
			case 0:
				playerScore += 1
			case 2:
				computerScore += 1
			}
		} else if playerValue == SCISSORS {
			switch computerValue {
			case 0:
				computerScore += 1
			case 1:
				playerScore += 1
			}
		}

		printCurrentScore(playerScore, computerScore)

		if playerScore == 3 {
			fmt.Println("Player wins")
			break
		} else if computerScore == 3 {
			fmt.Println("Computer wins")
			break
		}
	}

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

func printCurrentScore(playerScore int, computerScore int) {
	fmt.Println("Player score:", playerScore)
	fmt.Println("Computer score:", computerScore)
}
