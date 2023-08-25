package main

import "fmt"

func main() {
	ubCh := make(chan int)
	bCh := make(chan int, 5)

	// on unbuffered channel
	go func() {
		ubCh <- 1
		close(ubCh)
	}()

	for ub := range ubCh {
		fmt.Println("unbuffered:", ub)
	}

	// on buffered channel
	go func() { // you can assign value to bCh without creating a new gopher like this
		bCh <- 1
		bCh <- 5
		bCh <- 2
		bCh <- 3
		bCh <- 7
		close(bCh)
	}()

	for b := range bCh {
		fmt.Println("buffered:", b)
	}
}
