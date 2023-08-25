package main

import "fmt"

func main() {
	/**
	Buffered Channel
	*/
	bufferedChan := make(chan int) // max capacity: 1
	bufferedChan <- 1
	// bufferedChan <- 2 // will lead to deadlock error

	fmt.Println(<-bufferedChan) // 1

	/**
	Unbuffered Channel
	*/
	unbufferredChan := make(chan int, 2) // max capacity: 2
	unbufferredChan <- 1
	unbufferredChan <- 2 // still ok
	unbufferredChan <- 3 // will lead to deadlock error because exceeded max capacity

	fmt.Println(<-unbufferredChan) // 1
	fmt.Println(<-unbufferredChan) // 2
}
