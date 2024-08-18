package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	numGoroutines := 1000000 // Start with 1 million goroutines

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			select {} // Keeps the goroutine alive indefinitely
		}()
	}

	wg.Wait() // Wait for all goroutines to finish (in this case, they never will)
	fmt.Printf("Spawned %d goroutines\n", numGoroutines)

	// Print memory usage stats
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Memory used: %d KB\n", memStats.Alloc/1024)
}
