package main

import (
	"fmt"
)

func worker(ch <-chan int) {
	v := <-ch
	fmt.Println(v)
}

func main() {
	// Zombie goroutine - UNCOMMENT below and run to see
	// ch := make(chan int)
	// go worker(ch)

	// time.Sleep(time.Second)

	// Go memory model - UNCOMMENT below and run to see
	// data := "old value"

	// go func() {
	// 	data = "new value"
	// }()

	// go func() {
	// 	fmt.Println(data)
	// }()

	// time.Sleep(time.Second)
}
