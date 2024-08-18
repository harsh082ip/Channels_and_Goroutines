package main

import "fmt"

func main() {

Outerloop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("Breaking Inner loop, coz both i & j are 1 :/")
				break Outerloop
			}
			fmt.Printf("i: %v, j: %v \n", i, j)
		}
	}
}
