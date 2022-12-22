package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()
	fmt.Println("Type any character on the keyboard")

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			fmt.Println("Goodbye!")
			break
		}

		if key == keyboard.KeyEsc {
			fmt.Println("You have preesed esc key")
			break
		} else {
			fmt.Printf("You have pressed: %q\r", char)
		}

	}

}
