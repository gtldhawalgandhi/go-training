package main

import (
	"fmt"
	"sync"
	"time"
)

var global map[string]int = make(map[string]int)
var c = make(chan int, 4)

func thread1() {
	<-c // Grab the ticket
	global["s"] += 1
	fmt.Println("Grabbing thread1", global, time.Now())
	c <- 1 // Give it back
}

func thread2() {
	// If I read from an empty channel, then my code here gets blocked
	<-c
	global["s"] += 2
	fmt.Println("Grabbing thread2", global, time.Now())
	c <- 1
}

func mutexUsingChannels() {
	c <- 1 // Put the initial value into the channel
	go thread1()
	go thread2()

	time.Sleep(1 * time.Second)

	<-c
	fmt.Println("Channel Sync global value >> ", global["s"])
	c <- 1
}

func mutexUsingSync() {
	var mu = &sync.Mutex{}
	var mp = make(map[string]int)

	go func(mp map[string]int) {
		mu.Lock()
		fmt.Println(mp)

		mu.Unlock()

	}(mp)

	mu.Lock()
	mp["s"] = 99
	mu.Unlock()

	//time.Sleep(1 * time.Second)
}

func main() {
	//mutexUsingSync()
	mutexUsingChannels()
}
