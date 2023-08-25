package main

import (
	"fmt"
	"time"

)

func main() {
	ch := make(chan bool)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- true
	}()

	for {
		select {
		case <-ch:
			fmt.Println("channel has received something")
			return
		default:
			fmt.Println("waiting..")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
