package main

import (
	"fmt"
)

// Question 2: Buffered Channels
// Create a program where the main goroutine sends
// values into a buffered channel and a separate goroutine
//  receives and prints those values.

func Sender(ch chan int) {

	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func Receiver(ch chan int) {

	for i := range ch {
		fmt.Println(i)
	}
}

// func main() {

// 	ch := make(chan int, 10)

// 	go Sender(ch)
// 	go Receiver(ch)

// 	time.Sleep(time.Second * 2)

// }
