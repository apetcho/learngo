package level3

import (
	"fmt"
	"time"
)

// -*-------------*-
// -*- GoRoutine -*-
// -*-------------*-
func fn(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func GoRoutine() {
	fn("direct")
	go fn("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}

// -*------------*-
// -*- Channels -*-
// -*------------*-
func Channels() {
	fmt.Println()
	fmt.Println("-*------------*-")
	fmt.Println("-*- Channels -*-")
	fmt.Println("-*------------*-")
	messages := make(chan string)

	go func() { messages <- "ping" }()
	// -
	msg := <-messages
	fmt.Println(msg)
}

// -*--------------------*-
// -*- ChannelBuffering -*-
// -*--------------------*-
func ChannelBuffering() {
	fmt.Println()
	fmt.Println("-*----------------------*-")
	fmt.Println("-*- Channels buffering -*-")
	fmt.Println("-*----------------------*-")
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// -*--------------------------*-
// -*- ChannelSynchronization -*-
// -*--------------------------*-
func worker(done chan bool){
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func ChannelSynchronization(){
	fmt.Println()
	fmt.Println("-*---------------------------*-")
	fmt.Println("-*- Channels synchonization -*-")
	fmt.Println("-*---------------------------*-")

	done := make(chan bool, 1)
	go worker(done)
	<- done
}