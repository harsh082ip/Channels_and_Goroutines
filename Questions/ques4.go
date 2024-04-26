package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func Producer(ch chan int) {

// 	for i := 1; i <= 5; i++ {
// 		randomNum := rand.Intn(100)
// 		fmt.Println("Number sent: ", randomNum)
// 		ch <- randomNum
// 	}
// 	close(ch)

// }

// func Worker1(ch chan int) {
// 	for val := range ch {
// 		fmt.Println("From Worker 1: ", val)
// 	}

// }

// func Worker2(ch chan int) {
// 	for val := range ch {
// 		fmt.Println("From Worker 2: ", val)
// 	}

// }

// func Worker3(ch chan int) {
// 	for val := range ch {
// 		fmt.Println("From Worker 3: ", val)
// 	}

// }

// func main() {

// 	ch := make(chan int, 10)

// 	go Producer(ch)
// 	go Worker1(ch)
// 	go Worker2(ch)
// 	go Worker3(ch)

// 	time.Sleep(time.Second * 2)
// }

// In your code, each worker is reading from the same channel. Once a value is read from the channel by one worker, it's removed from the channel and not available for the other workers to read. This behavior is because channels in Go are designed for communication between goroutines and each value sent through a channel is consumed exactly once.

// If you want each worker to receive the same value from the producer, you should either:

// Broadcast the value to multiple channels, one for each worker.
// Have each worker perform the same work independently.

func Producer(ch1, ch2, ch3 chan int) {
	for i := 1; i <= 5; i++ {
		randomNum := rand.Intn(100)
		fmt.Println("Number sent: ", randomNum)
		ch1 <- randomNum
		ch2 <- randomNum
		ch3 <- randomNum
	}
	close(ch1)
	close(ch2)
	close(ch3)
}

func Worker1(ch chan int) {
	for val := range ch {
		fmt.Println("From Worker 1:", val)
	}
}

func Worker2(ch chan int) {
	for val := range ch {
		fmt.Println("From Worker 2:", val)
	}
}

func Worker3(ch chan int) {
	for val := range ch {
		fmt.Println("From Worker 3:", val)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go Producer(ch1, ch2, ch3)
	go Worker1(ch1)
	go Worker2(ch2)
	go Worker3(ch3)

	time.Sleep(time.Second * 2)
}
