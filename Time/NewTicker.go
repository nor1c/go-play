package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second) // run every seconds
	defer ticker.Stop()

	done := make(chan bool)

	// timeout after 10 seconds
	go func() {
		time.Sleep(time.Second * 5)
		done <- true
	}()

Stop:
	for {
		select {
		case <-done:
			fmt.Println("done!") // should be done and stopped the ticker
			break Stop
		case t := <-ticker.C:
			fmt.Println("Current time:", t)
		}
	}
}
