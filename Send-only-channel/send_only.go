package main

import (
	"fmt"
)

func sendData(data chan<- int) {
	// Send some data through the channel
	data <- 42
}

func sendData2(data chan int) {
	// Send some data through the channel
	data <- 42
}

func main() {
	// Create a channel for sending data
	ch := make(chan int)

	// Pass the channel to the function that only sends data
	sendData(ch)
	// sendData2(ch)

	// Trying to receive data from the channel will result in a compilation error
	val := <-ch

	fmt.Println("Data sent successfully!", val)
}
