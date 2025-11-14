package main

import (
	"fmt"
	"time"
)

func worker(i int) {
	fmt.Printf("wroker %d started.\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("wroker %d finished.\n", i)
}

func main() {
	limit := 3

	sem := make(chan struct{}, limit)

	for i := 0; i < 10; i++ {
		sem <- struct{}{}
		go func(j int) {
			worker(j)
			<-sem
		}(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("All finished")

}
