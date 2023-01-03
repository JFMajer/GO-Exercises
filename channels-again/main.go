package main

import "time"

var chan1 = make(chan string)
var chan2 = make(chan string)

func task1() {
	time.Sleep(time.Second)
	chan1 <- "one"
}

func task2() {
	time.Sleep(time.Second * 2)
	chan2 <- "two"
}

func main() {
	go task1()
	go task2()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			println("received", msg1)
		case msg2 := <-chan2:
			println("received", msg2)
		}
	}

}
