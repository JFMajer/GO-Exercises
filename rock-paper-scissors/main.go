package main

import (
	"myapp/game"
)

func main() {

	displayChan := make(chan string)
	roundChan := make(chan int)

	game := game.Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}

	game.ClearScreen()
	go game.Rounds()
	game.PrintIntro()

	for {
		game.RoundChan <- 1
		<-game.RoundChan

		if game.Round.RoundNumber > 3 {
			break
		}

		if !game.PlayRound() {
			game.RoundChan <- -1
			<-game.RoundChan
		}
		game.PrintCurrentScore()
	}
	game.PrintGameSummary()
}
