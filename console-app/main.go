package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	fmt.Println("Type any character on the keyboard")

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			fmt.Printf("You pressed %c\n", char)
		} else {

		}

		if key == keyboard.KeyEsc {
			fmt.Println("You have preesed esc key")
			break
		}

	}

}
