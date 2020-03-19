package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// Add number of goroutines
	wg.Add(2)

	go sleep(&wg, time.Second*1)
	go sleep(&wg, time.Second*2)

	// wait used to block all goroutines till the done method called
	wg.Wait()
	fmt.Println("All go routines finished")
}

func sleep(wg *sync.WaitGroup, t time.Duration) {
	// called by each go routines finished execution
	defer wg.Done()
	time.Sleep(t)
	fmt.Println("Finished goroutine exceution")
}
