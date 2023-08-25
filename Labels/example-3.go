package main

import (
	"log"
	"time"

)

func main() {
	tickRate := 1 * time.Second

	stopper := time.After(6 * time.Second)

	ticker := time.NewTicker(tickRate)

loop:
	for {
		select {
		case <-ticker.C:
			log.Println("tick")
		case <-stopper:
			break loop
		}
	}

	log.Println("all process finished!")
}
