package main

import (
	"fmt"
	"sync"
	"time"
)

// Question 3: Channel Synchronization
// Write a Go program where two goroutines are concurrently
// incrementing a shared variable. Use channels to synchronize
//  access to the variable and ensure both goroutines update it correctly.

var value int = 0
var mu sync.Mutex

func Increment1() {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Increment1...")

	for i := 1; i <= 5; i++ {
		value++
	}
}

func Increment2() {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println("Increment2...")

	for i := 1; i <= 5; i++ {
		value++
	}
}

func main() {

	go Increment1()
	go Increment2()

	time.Sleep(time.Second * 3)

	fmt.Println(value)
}
