package main

import "fmt"

// send-only "chan<-"
func producer(v int, ch chan<- int) {
	ch <- v
}

// receive-only "<-chan"
func consumer(ch <-chan int) int {
	return <-ch
}

func main() {
	ch := make(chan int)
	go producer(42, ch)
	res := consumer(ch) // same as: res := <-ch
	fmt.Println("received", res)
}
