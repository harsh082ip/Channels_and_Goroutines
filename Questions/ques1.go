package main

// import "fmt"

// Question 1: Channel Basics
// Write a Go program that creates a channel, sends integers
// from 1 to 10 into the channel using a goroutine,
// and then prints them out from the main goroutine.

func Printvalues(ch chan int) {

	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// func main() {

// 	ch := make(chan int, 10)
// 	go Printvalues(ch)
// 	for p := range ch {
// 		fmt.Println(p)
// 	}
// }
