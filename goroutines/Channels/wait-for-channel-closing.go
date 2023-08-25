package main

import (
	"fmt"
)

func producer(v int, ch chan<- int) {
	defer close(ch)
	ch <- v
	fmt.Println("ch receives a value")
}

func main() {
	nc := make(chan int, 2)

	go producer(5, nc)

	<-nc // wait until channel `ch` receives a value before processing to the next line

	fmt.Println("done!")
}
