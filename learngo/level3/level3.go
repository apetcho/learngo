package level3

import (
	"fmt"
	"sync"
	"sync/atomic"
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

// -*------------*-
// -*- Timeouts -*-
// -*------------*-
func Timeouts() {
	fmt.Println()
	fmt.Println("-*------------*-")
	fmt.Println("-*- Timeouts -*-")
	fmt.Println("-*------------*-")

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

// -*----------------------*-
// -*- NonBlockingChannel -*-
// -*----------------------*-
func NonBlockingChannel() {
	fmt.Println()
	fmt.Println("-*-----------------------------------*-")
	fmt.Println("-*- Non-blocking channel operations -*-")
	fmt.Println("-*-----------------------------------*-")
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

// -*-------------------*-
// -*- ClosingChannels -*-
// -*-------------------*-
func ClosingChannels() {
	fmt.Println()
	fmt.Println("-*--------------------*-")
	fmt.Println("-*- Closing Channels -*-")
	fmt.Println("-*--------------------*-")
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}

// -*-------------------*-
// -*- ChannelIterator -*-
// -*-------------------*-
func ChannelIterator() {
	fmt.Println()
	fmt.Println("-*-----------------------*-")
	fmt.Println("-*- Range over channels -*-")
	fmt.Println("-*-----------------------*-")

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}

// -*----------*-
// -*- Timers -*-
// -*----------*-
func Timers() {
	fmt.Println()
	fmt.Println("-*----------*-")
	fmt.Println("-*- Timers -*-")
	fmt.Println("-*----------*-")

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

// -*-----------*-
// -*- Tickers -*-
// -*-----------*-
func Tickers() {
	fmt.Println()
	fmt.Println("-*-----------*-")
	fmt.Println("-*- Tickers -*-")
	fmt.Println("-*-----------*-")

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

// -*---------------*-
// -*- WorkerPools -*-
// -*---------------*-
func myworker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func WorkerPools() {
	fmt.Println()
	fmt.Println("-*----------------*-")
	fmt.Println("-*- Worker Pools -*-")
	fmt.Println("-*----------------*-")
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go myworker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		fmt.Println("result =", <-results)
	}
}

// -*--------------*-
// -*- WaitGroups -*-
// -*--------------*-
func waitGroupWorker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroups() {
	fmt.Println()
	fmt.Println("-*--------------*-")
	fmt.Println("-*- WaitGroups -*-")
	fmt.Println("-*--------------*-")
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			waitGroupWorker(i)
		}()
	}
	wg.Wait()
}

// -*----------------*-
// -*- RateLimiting -*-
// -*----------------*-
func RateLimiting() {
	fmt.Println()
	fmt.Println("-*-----------------*-")
	fmt.Println("-*- Rate Limiting -*-")
	fmt.Println("-*-----------------*-")
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("bursty request", req, time.Now())
	}
}

// -*------------------*-
// -*- AtomicCounters -*-
// -*------------------*-
func AtomicCounters() {
	fmt.Println()
	fmt.Println("-*-------------------*-")
	fmt.Println("-*- Atomic Counters -*-")
	fmt.Println("-*-------------------*-")
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops.Load())
}

// -*-----------*-
// -*- Mutexes -*-
// -*-----------*-
type Container struct {
	mtx      sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.counters[name]++
}

func Mutexes() {
	fmt.Println()
	fmt.Println("-*-----------*-")
	fmt.Println("-*- Mutexes -*-")
	fmt.Println("-*-----------*-")
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
