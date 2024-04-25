package main

import (
	"fmt"
	"time"
)

func PrintSomething(c chan string) {
	time.Sleep(time.Second * 5)
	c <- "Hello"
}

func PrintSomething2(c chan string) {
	time.Sleep(time.Second * 10)
	c <- "Hello 10"
}

func PrintSomething3(c chan string) {
	time.Sleep(time.Second * 12)
	c <- "break"
}

func main() {

	data := make(chan string)

	go PrintSomething2(data)
	go PrintSomething(data)
	go PrintSomething3(data)

	for {
		msg := <-data
		if msg == "break" {
			fmt.Println("break triggered...")
			break

		}
		fmt.Println(msg)

	}

	// msg := <-data
	// fmt.Println(msg)

	// m := <-data
	// fmt.Println(m)

	fmt.Println("Nothing...")
	// os.Exit(1)

}
