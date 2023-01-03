package game

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

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	// use select to process input
	for {
		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round
			g.RoundChan <- 1
		case msg := <-g.DisplayChan:
			fmt.Println(msg)
		}

	}
}

func (g *Game) ClearScreen() {
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

func (g *Game) PrintIntro() {
	g.DisplayChan <- fmt.Sprintf(`
Rock, Paper, Scissors
---------------------
Game is played for three rounds, the player with the most wins is the winner, good luck!
	`)
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	playerValue := -1

	fmt.Printf("Round %d\n", g.Round.RoundNumber)
	fmt.Print("Please enter rock, paper or scissors -> ")

	playerChoice, _ := reader.ReadString('\n')
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

	g.DisplayChan <- fmt.Sprintf("Player picked %d which is %s, computer picked %d which means %s\n", playerValue, playerChoice, computerValue, computerChoice)

	if playerValue == computerValue {
		g.DisplayChan <- "It's a tie"
		return false
	} else if playerValue == ROCK {
		switch computerValue {
		case 1:
			g.Round.ComputerScore += 1
			g.DisplayChan <- "Computer wins"
		case 2:
			g.Round.PlayerScore += 1
			g.DisplayChan <- "Player wins"
		}
		return true
	} else if playerValue == PAPER {
		switch computerValue {
		case 0:
			g.Round.PlayerScore += 1
			g.DisplayChan <- "Player wins"
		case 2:
			g.Round.ComputerScore += 1
			g.DisplayChan <- "Computer wins"
		}
		return true
	} else if playerValue == SCISSORS {
		switch computerValue {
		case 0:
			g.Round.ComputerScore += 1
			g.DisplayChan <- "Computer wins"
		case 1:
			g.Round.PlayerScore += 1
			g.DisplayChan <- "Player wins"
		}
		return true
	}
	return false
}

func (g *Game) PrintCurrentScore() {
	g.DisplayChan <- fmt.Sprintf("Player: %d, Computer: %d", g.Round.PlayerScore, g.Round.ComputerScore)
	g.DisplayChan <- "---------------------------------------"
}

func (g *Game) PrintGameSummary() {
	g.DisplayChan <- "Game Summary"
	g.DisplayChan <- fmt.Sprintf("Player: %d, Computer: %d", g.Round.PlayerScore, g.Round.ComputerScore)
	g.DisplayChan <- "---------------------------------------"
}
