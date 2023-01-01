package main

import (
	"bufio"
	"fmt"
	"myapp/myLogger"
	"os"
	"time"
)

func main() {
	// read input from the user 5 times and write it to a log

	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	go myLogger.ListenForLog(ch)

	for i := 0; i < 5; i++ {
		fmt.Println("Enter something: ")
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}

}
