// package main

// import (
// 	"fmt"
// )

// func channel_num(number chan int) {
// 	//empty function
// }

// func channel_text(message chan string) {
// 	//empty function
// }

// func main() {
// 	num := make(chan int)     //initializing the channel
// 	text := make(chan string) //initializing the channel

// 	go channel_num(num)   //calling as a goroutine
// 	go channel_text(text) //calling as a goroutine

// 	select {
// 	case channel_1 := <-num: //channel1 is receiving data from channel num
// 		fmt.Println("Data:", channel_1)

// 	case channel_2 := <-text: //channel2 is receiving data from channel text
// 		fmt.Println("Data:", channel_2)

// 		// default:
// 		// 	fmt.Println("End..")
// 	}

// }

package main

import (
	"fmt"
)

func channel_num(number chan int) {

	number <- 239 //function with a data

}

func main() {
	num := make(chan int) //initializing the channel

	go channel_num(num) //calling as a goroutine

	i := <-num //sending data from channel num
	j := <-num //sending data from channel num
	fmt.Println("Value of Channel i,j =", i, j)
}
