package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var wg sync.WaitGroup

func main() {

	t := time.Now()
	nums := []int{1, 2, 3, 4, 5, 6}
	var finalnum []int

	for _, num := range nums {
		wg.Add(1)
		go customAppend(&finalnum, num)
	}

	wg.Wait()

	fmt.Println(finalnum)
	fmt.Println("Time taken: ", time.Since(t))
}

func customAppend(finalnum *[]int, num int) {
	time.Sleep(time.Second)

	mu.Lock()
	*finalnum = append(*finalnum, num*2)
	mu.Unlock()

	wg.Done()
}
