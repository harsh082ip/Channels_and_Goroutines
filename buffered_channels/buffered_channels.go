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

package main

import (
	"fmt"
)

func timesThree(arr []int, ch chan int) {
	for _, elem := range arr {
		ch <- elem * 3
	}
}

func main() {
	fmt.Println("We are executing a goroutine")
	arr := []int{2, 3, 4}
	ch := make(chan int, len(arr))
	go timesThree(arr, ch)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Result: %v \n", <-ch)
	}
}
