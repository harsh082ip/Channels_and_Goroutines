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

// package main

// import (
// 	"fmt"
// )

// func channel_num(number chan int) {

// 	number <- 239 //function with a data

// }

// func main() {
// 	num := make(chan int) //initializing the channel

// 	go channel_num(num) //calling as a goroutine

// 	i := <-num //sending data from channel num
// 	j := <-num //sending data from channel num
// 	fmt.Println("Value of Channel i,j =", i, j)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func Receiver(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for {
// 		select {
// 		case val := <-ch:
// 			// for i := 1; i <= 1000; i++ {
// 			// 	fmt.Println(val, i)
// 			// }
// 			// time.Sleep(time.Second * 3)
// 			fmt.Println(val)
// 		}
// 	}
// }

// func Reciver2(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for val := range ch {
// 		fmt.Println(val)
// 	}

// }

// func main() {

// 	var wg sync.WaitGroup

// 	ch := make(chan int, 2200)

// 	// go Receiver(ch, &wg)
// 	for i := 1; i <= 2000; i++ {
// 		ch <- i
// 	}
// 	close(ch)

// 	time.Sleep(time.Second * 4)

// 	wg.Add(1)
// 	go Reciver2(ch, &wg)
// 	wg.Wait()

// }

// Closing the channel simply signals to the receiving goroutine
// that no more values will be sent on the channel, but it doesn't
// prevent the receiver from processing the values that are already
// in the channel. Therefore, the Receiver2 goroutine is still able
// to receive and process the values sent before the channel was
// closed.

package main

import "fmt"

func Receiver(ch1, ch2 chan int) {

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(val2)
	}
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go Receiver(ch1, ch2)
	ch1 <- 2
	ch1 <- 1

}
