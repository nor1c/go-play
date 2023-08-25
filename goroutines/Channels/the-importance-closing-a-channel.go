package main

import (
	"log"

)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch) // without this, it will return a deadlock error, means that line 15 is still expecting values from channel `ch`
	}()

	for n := range ch {
		log.Println("n:", n)
	}
}
