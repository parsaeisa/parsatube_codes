package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println("New item from the channel: ", val)
	}

	fmt.Println("The channel is closed")
}
