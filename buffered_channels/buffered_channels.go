// package main

// import (
// 	"fmt"
// 	"time"
// )

// func timesThree(arr []int, ch chan int) {
// 	for _, elem := range arr {
// 		fmt.Println(elem * 3)
// 		ch <- elem * 3
// 	}
// }

// func main() {
// 	fmt.Println("We are executing a goroutine")
// 	arr := []int{2, 3, 4}
// 	ch := make(chan int)
// 	go timesThree(arr, ch)
// 	time.Sleep(time.Second) // we let the goroutine return all the values
// 	result := <-ch
// 	fmt.Printf("The result is: %v", result)
// }

// package main

// import (
// 	"fmt"
// )

// func timesThree(arr []int, ch chan int) {
// 	for _, elem := range arr {
// 		ch <- elem * 3
// 	}
// }

// func main() {
// 	fmt.Println("We are executing a goroutine")
// 	arr := []int{2, 3, 4}
// 	ch := make(chan int, len(arr))
// 	go timesThree(arr, ch)
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Printf("Result: %v \n", <-ch)
// 	}
// }

package main

import (
	"fmt"
	"time"
)

// func main() {
// 	// Creating a buffered channel with a capacity of 3
// 	channel := make(chan int)

// 	// Producer goroutine
// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			channel <- i // Send number to channel
// 			fmt.Println("Produced:", i)
// 		}
// 		close(channel) // Close the channel when done producing
// 	}()

// 	// Consumer goroutine
// 	go func() {
// 		for num := range channel {
// 			fmt.Println("Consumed:", num)
// 			time.Sleep(time.Second) // Simulate time-consuming task
// 		}
// 	}()

// 	// Wait for a moment before exiting
// 	time.Sleep(6 * time.Second)
// }

func main() {

	ch := make(chan int)

	go func(ch chan int) {
		time.Sleep(time.Second * 5)
		val := <-ch
		fmt.Println(val)
	}(ch)

	ch <- 14
	fmt.Println("End...")
	time.Sleep(time.Second * 6)
}
