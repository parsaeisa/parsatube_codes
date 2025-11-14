package main

import (
	"fmt"
	"time"
)

func firstBufferedChannel() {
	ch := make(chan int, 3) // unbuffered

	start := time.Now()

	for i := 1; i <= 3; i++ {
		go func(n int) {
			fmt.Println("Sending:", n)
			ch <- n // blocks until main receives
			fmt.Println("Sent:", n)
		}(i)
	}

	time.Sleep(2 * time.Second)

	for i := 1; i <= 3; i++ {
		val := <-ch
		fmt.Println("Received:", val)
	}

	fmt.Println("Total time:", time.Since(start))
}
