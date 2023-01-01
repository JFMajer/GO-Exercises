package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 0; i < 200; i++ {
		fmt.Print("a")
	}
}

func b() {
	for i := 0; i < 200; i++ {
		fmt.Print("b")
	}
}

func main() {
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("End main")
}
