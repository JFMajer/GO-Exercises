package main

import (
	"fmt"
	"log"
	"strconv"

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

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocca"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("===============")

	for k, v := range coffees {
		fmt.Println(k, v)
	}

	for {
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			fmt.Println("Goodbye!")
			break
		}

		i, _ := strconv.Atoi(string(char))
		t := fmt.Sprintf("You chose %s", coffees[i])
		fmt.Println(t)

		if key == keyboard.KeyEsc {
			fmt.Println("You have preesed esc key")
			break
		} else {
			fmt.Printf("You have pressed: %q\r", char)
		}

	}

}
