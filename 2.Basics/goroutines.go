package main

import (
	"fmt"
	"time"
)

// Go routine >> main go routine
func main() {

	// pipes Send > =============== > Receive
	//ch := make(chan int, 1) // Buffered Channel
	ch := make(chan int)

	go func() {

		fmt.Println(">> >> ", <-ch)
	}()

	// Non Buffered Channel is going to block this line
	ch <- 567

	// Buffered Channel we can use this method
	//fmt.Println(">> >> ", <-ch)

	ch2 := make(chan int, 3)

	go func() {
		ch2 <- 22
		ch2 <- 23
		ch2 <- 24

		close(ch2)
	}()

	//fmt.Println(<-ch2)
	//fmt.Println(<-ch2)
	//fmt.Println(<-ch2)

	//for i := 0; i < 3; i++ {

	//fmt.Println(<-ch2)

	//}

	// Our for loop will exit when our channel is closed
	// blocking
	for v := range ch2 {
		fmt.Println(v)
	}

	ch3 := make(chan string)

	// Go routine runtime scheduler
	go func() {
		//time.Sleep(10 * time.Second)
		ch3 <- "Hello From Go Routine"
	}()

	// Les pause the execution for 100 ms
	//time.Sleep(100 * time.Microsecond)

	timeout := time.After(100 * time.Millisecond)

	for {
		select {
		// Receive operation
		case v := <-ch3:
			fmt.Println("From select >> ", v)
		case <-timeout:
			fmt.Println("Timing out, too slow")
			return
			// If you dont have default then this select will get blocked
			//default:
			//fmt.Println("No operation found")
		}
	}
}
