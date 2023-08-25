package main

import "fmt"

func fib(c chan int, done chan bool) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-done:
			fmt.Println("done")
			return
		}
	}
}

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		done <- true
	}()

	fib(c, done)
}
