package main

import (
	"fmt"
	"sync"
)

// go run -race waitGroup.go

// Refer Go routines if you have not already
func main() {
	var wg = &sync.WaitGroup{}
	var mu = &sync.Mutex{}

	// maps in Go are not thread safe and are subjected to race condition
	data := make(map[string]int)

	// We add number of go routines, that we are going to create, into "Add" method

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.Mutex, idx int) {
			mu.Lock()
			defer mu.Unlock()
			data["a"] += idx

			wg.Done()
		}(wg, mu, i)

	}

	// wait until we get the data (this will block our code from executing any further) so make sure we call wg.Done() from our goroutine
	wg.Wait()
	//time.Sleep(1 * time.Second)

	mu.Lock()
	fmt.Println(data)
	mu.Unlock()
}
