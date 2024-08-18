package main

import (
	"fmt"
	"time"
)

func Sender(ch []chan int) {

	for i := 1; i <= 6; i++ { // values
		for j := 0; j < 4; j++ { // broadcast
			ch[j] <- i
		}
	}
}

func Worker1(ch chan int) {
	defer func() {
		fmt.Println("Worker 1 closed")
		close(ch)
	}()
	for {
		select {
		case val := <-ch:
			if val == 6 {
				return
			}
			fmt.Println("1 time", val)

		}
	}
}

func Worker2(ch chan int) {
	defer func() {
		fmt.Println("Worker 2 closed")
		close(ch)
	}()
	for {
		select {
		case val := <-ch:
			if val == 6 {
				return
			}
			fmt.Println("two times", val*val)
		}
	}
}

func Worker3(ch chan int) {
	defer func() {
		fmt.Println("Worker 3 closed")
		close(ch)
	}()
	for {
		select {
		case val := <-ch:
			if val == 6 {
				return
			}
			fmt.Println("Three times", val*val*val)
		}
	}
}

func Worker4(ch chan int) {
	defer func() {
		fmt.Println("Worker 4 closed")
		close(ch)
	}()
	for {
		select {
		case val := <-ch:
			if val == 6 {
				return
			}
			fmt.Println("Fourtimes", val*val*val*val)
		}
	}
}

func main() {

	channels := [4]chan int{}

	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
	}

	go Worker1(channels[0])
	go Worker2(channels[1])
	go Worker3(channels[2])
	go Worker4(channels[3])

	go Sender(channels[:])

	time.Sleep(time.Second * 10)

}
