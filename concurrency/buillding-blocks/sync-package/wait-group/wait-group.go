package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // <1>
	go func() {
		defer wg.Done() // <2>
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1) // <1>
	go func() {
		defer wg.Done() // <2>
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2)
	}()

	wg.Wait() // <3>
	fmt.Println("All goroutines complete.")

	const numGreeters = 5
	var bulkWg sync.WaitGroup

	bulkWg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go bulkGoroutine(&bulkWg, i)
	}

	bulkWg.Wait()
}

func bulkGoroutine(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("\nHello from id %v", id)
}
