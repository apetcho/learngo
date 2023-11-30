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
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func ChannelSynchronization() {
	fmt.Println()
	fmt.Println("-*---------------------------*-")
	fmt.Println("-*- Channels synchonization -*-")
	fmt.Println("-*---------------------------*-")

	done := make(chan bool, 1)
	go worker(done)
	<-done
}

// -*---------------------*-
// -*- ChannelDirections -*-
// -*---------------------*-
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func ChannelDirections() {
	fmt.Println()
	fmt.Println("-*-----------------------*-")
	fmt.Println("-*- Channels Directions -*-")
	fmt.Println("-*-----------------------*-")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// -*----------*-
// -*- Select -*-
// -*----------*-
func Select() {
	fmt.Println()
	fmt.Println("-*----------*-")
	fmt.Println("-*- Select -*-")
	fmt.Println("-*----------*-")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
